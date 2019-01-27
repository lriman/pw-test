package models

type Response struct {
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}


type ResponseProfile struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string   `json:"email"`
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
	Amount    uint  `json:"amount"`
}
