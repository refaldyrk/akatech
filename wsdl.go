package main

import (
	"akatech/dto"
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func callSOAPService(userID string) (*dto.UserXMLResponse, error) {
	request := dto.UserXMLResponse{
		UserID: userID,
	}

	requestXML, err := xml.MarshalIndent(request, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %v", err)
	}

	soapEnvelope := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:web="http://example.com/soap-service">
   <soapenv:Header/>
   <soapenv:Body>
` + string(requestXML) + `
   </soapenv:Body>
</soapenv:Envelope>`)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://example.com/soap-service", bytes.NewBuffer(soapEnvelope))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("SOAPAction", "GetUserDetails")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	var soapResponse dto.UserXMLResponse
	err = xml.Unmarshal(body, &soapResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	return &soapResponse, nil
}
