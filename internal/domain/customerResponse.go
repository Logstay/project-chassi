package domain

type CustomerResponse struct {
	Code     int         `json:"code"`
	Response interface{} `json:"response"`
}
