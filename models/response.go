package models

import (
	"encoding/json"
)

// Response Struct to return the response of the Rede's API
type Response struct {
	// Return code
	ReturnCode string `json:"returnCode"`
	// Return message
	ReturnMessage string `json:"returnMessage"`
	// Reference to the Request.Reference code
	Reference string `json:"reference"`
	// Unique payment indentifier
	Tid string `json:"tid"`
	// Sequencial number returned by Rede
	Nsu string `json:"nsu"`
	// Authorization code by the card's manufactory
	AuthorizationCode string `json:"authorizationCode"`
	// DateTime time when rede processed the request
	DateTime string `json:"dateTime"`
	// Amount R$10,00 equals 1000
	Amount int `json:"amount"`
	// CardBin first 6 digits of the credit card
	CardBin string `json:"cardBin"`
	// Last4 last four digits of the credit card
	Last4 string `json:"last4"`
	// BrandTid
	BrandTid string `json:"brandTid,omitempty"`
	// Brand are all possible data returned by the Rede to new Format
	Brand brandResponse `json:"brand,omitempty"`
	// Links are all possible links returned by the Rede (refund for example)
	Links []links `json:"links"`
}

type brandResponse struct {
	Name              string `json:"name,omitempty"`
	ReturnCode        string `json:"returnCode,omitempty"`
	ReturnMessage     string `json:"returnMessage,omitempty"`
	AuthorizationCode string `json:"authorizationCode,omitempty"`
	BrandTid          string `json:"brandTid,omitempty"`
}

// links all links returned by Rede
type links struct {
	Method string `json:"method"`
	Rel    string `json:"rel"`
	Href   string `json:"href"`
}

// String return a string containing all data about the response
func (r Response) String() string {
	msg, _ := json.MarshalIndent(r, "", " ")
	return string(msg) //fmt.Sprintf("Return Code: %s\nReturn Message: %s\nReference: %s\nTid: %s\nNsu: %s\nAuthorizationCode: %s\nDateTime: %s\nAmount: %d\nCardBin: %s\nLast4: %s\nBrandTid: %s\nLinks: %s\n", r.ReturnCode, r.ReturnMessage, r.ReturnMessage, r.Tid, r.Nsu, r.AuthorizationCode, r.DateTime, r.Amount, r.CardBin, r.Last4, r.BrandTid, r.Links)
}
