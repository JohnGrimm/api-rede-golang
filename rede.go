package rede

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	errors2 "github.com/JohnGrimm/api-rede-golang/apierrs"
	"github.com/JohnGrimm/api-rede-golang/utils"

	"github.com/JohnGrimm/api-rede-golang/login"
	"github.com/JohnGrimm/api-rede-golang/models"
)

// Rede interface for the Rede's API
type Rede interface {
	Pay(p *models.Payment) (*models.Response, error)
	TestCard(p *models.Payment) (*models.Response, error)
	Capture(tid string) (*models.Response, error)
}

type rede struct {
	config *login.Login
}

// NewRede instantiate a new Rede API object
func NewRede(pv string, ik string, isProduction bool) Rede {
	return &rede{
		config: &login.Login{
			PV:             pv,
			IntegrationKey: ik,
			IsProduction:   isProduction,
		},
	}
}

// Pay a method to do the payment
func (r rede) Pay(req *models.Payment) (*models.Response, error) {
	postParameters, err := req.ToJSON()
	if err != nil {
		return nil, errors2.APIErr(err.Error())
	}

	body := ""
	if r.config.IsProduction {
		_, _, body = doPostRequest(utils.APIBaseURL(), "POST", postParameters, r.config)

	} else {
		_, _, body = doPostRequest(utils.APIBaseURLTest(), "POST", postParameters, r.config)

	}

	var parseHeader models.Response
	err = json.Unmarshal([]byte(body), &parseHeader)
	if err != nil {
		return nil, errors2.APIErr(err.Error())
	}

	if strings.Compare(parseHeader.ReturnCode, "00") != 0 && strings.Compare(parseHeader.ReturnCode, "174") != 0 {
		err = errors2.APIErr("The payment was not successful!")
	}

	return &parseHeader, err
}

// Capture a method to do the payment
func (r rede) Capture(tid string) (*models.Response, error) {

	body := ""
	if r.config.IsProduction {
		_, _, body = doPostRequest(utils.APIBaseURL()+tid, "PUT", []byte(""), r.config)

	} else {
		_, _, body = doPostRequest(utils.APIBaseURLTest()+tid, "PUT", []byte(""), r.config)

	}

	var parseHeader models.Response
	err := json.Unmarshal([]byte(body), &parseHeader)
	if err != nil {
		return nil, errors2.APIErr(err.Error())
	}

	if strings.Compare(parseHeader.ReturnCode, "00") != 0 && strings.Compare(parseHeader.ReturnCode, "174") != 0 {
		err = errors2.APIErr("The payment was not successful!")
	}

	return &parseHeader, err
}

// TestCard is a test function to see if the card is valid
func (r rede) TestCard(req *models.Payment) (*models.Response, error) {
	req.Amount = 0

	payment, err := r.Pay(req)
	if err != nil {
		err = errors2.APIErr("The card is not valid!")
	}

	return payment, err
}

// doPostRequest do the low level needs for the requests
func doPostRequest(url string, method string, content []byte, login *login.Login) (string, string, string) {

	var jsonStr = content
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	req.SetBasicAuth(login.PV, login.IntegrationKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Transaction-Response", "brand-return-opened")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	return resp.Status, fmt.Sprint(resp.Header), string(body)
}
