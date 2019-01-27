package models

import "time"

type Response struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

type ResponseProfile struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Balance uint   `json:"balance"`
}

type ResponseAutoComplete struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type RequestSignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestSignUp struct {
	Email     string `json:"email"`
	Name      string `json:"name"`
	Password1 string `json:"password1"`
	Password2 string `json:"password2"`
}

type RequestAutoComplete struct {
	Name string `json:"name"`
}

type RequestTransfer struct {
	Recipient string `json:"recipient"`
	Amount    uint   `json:"amount"`
}

type RequestHistory struct {
	TimestampMin *int64  `json:"timestampMin"`
	TimestampMax *int64  `json:"timestampMax"`
	AmountMin    *uint   `json:"amountMin"`
	AmountMax    *uint   `json:"amountMax"`
	Name         string  `json:"name"`
	Sort         string  `json:"sort"`
	Size         uint    `json:"size"`
	Offset       uint    `json:"offset"`
}

func (r *RequestHistory)DateMin() *time.Time{
	if r.TimestampMin != nil {
		t := time.Unix(*r.TimestampMin, 0)
		return &t
	}
	return nil
}

func (r *RequestHistory)DateMax() *time.Time{
	if r.TimestampMax != nil {
		t := time.Unix(*r.TimestampMax, 0)
		return &t
	}
	return nil
}
