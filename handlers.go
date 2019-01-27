package main

import (
	"encoding/json"
	"net/http"
	"log"
	"github.com/pw-test/models"
)


func (a *App) SignInHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("SignIn")

	req := models.RequestSignIn{}
	resp := models.Response{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp.Error = err.Error()
		json.NewEncoder(w).Encode(resp)
		log.Println("Response for SignIn:", resp)
		return
	}

	u, err := a.SignInController(req.Email, req.Password)
	if err != nil {
		resp.Error = err.Error()
	} else {
		resp.Data = u
	}

	log.Println("Response for SignIn:", resp)
	json.NewEncoder(w).Encode(resp)
}


func (a *App) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("SignUp")

	req := models.RequestSignUp{}
	resp := models.Response{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp.Error = err.Error()
		json.NewEncoder(w).Encode(resp)
		log.Println("Response for SignUp:", resp)
		return
	}

	err = a.SignUpController(req.Email, req.Password1, req.Password2, req.Name)
	if err != nil {
		resp.Error = err.Error()
	}

	log.Println("Response for SignUp:", resp)
	json.NewEncoder(w).Encode(resp)
}


func (a *App) ProfileHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request for Profile")

	resp := models.Response{}

	profile, err := a.ProfileController(r.Context().Value("user").(string))
	if err != nil {
		resp.Error = err.Error()
	} else {
		resp.Data = profile
	}

	log.Println("Response for Profile:", resp)
	json.NewEncoder(w).Encode(resp)
}


func (a *App) AutoCompleteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("AutoCompleteHandler")

	req := models.RequestAutoComplete{}
	resp := models.Response{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp.Error = err.Error()
		json.NewEncoder(w).Encode(resp)
		log.Println("Response for AutoCompleteHandler:", resp)
		return
	}

	users, err := a.AutoCompleteController(req.Name)
	if err != nil {
		resp.Error = err.Error()
	} else {
		resp.Data = users
	}

	log.Println("Response for AutoCompleteHandler:", resp)
	json.NewEncoder(w).Encode(resp)
}

func (a *App) TransferHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("TransferHandler")

	req := models.RequestTransfer{}
	resp := models.Response{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp.Error = err.Error()
		json.NewEncoder(w).Encode(resp)
		log.Println("Response for TransferHandler:", resp)
		return
	}

	err = a.TransferController(r.Context().Value("user").(string), req.Recipient, req.Amount)
	if err != nil {
		resp.Error = err.Error()
	}

	log.Println("Response for TransferHandler:", resp)
	json.NewEncoder(w).Encode(resp)
}

