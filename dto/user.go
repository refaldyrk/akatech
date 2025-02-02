package dto

import "encoding/xml"

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserXMLResponse struct {
	XMLName xml.Name `xml:"GetUserDetailsResponse"`
	UserID  string   `xml:"user_id"`
	Name    string   `xml:"name"`
	Email   string   `xml:"email"`
}
