package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"

	"github.com/pw-test/models"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize(host, user, pwd, db string) {
	var err error

	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, pwd, db)
	a.DB, err = gorm.Open("postgres", conn)

	if err != nil {
		log.Fatal("Could not connect to Database:", err)
	}

	a.DB.AutoMigrate(models.User{})
	a.DB.AutoMigrate(models.Transaction{})

	a.Router = mux.NewRouter()
	a.Router.Use(JwtAuthentication)

	a.Router.HandleFunc("/api/sign-in", a.SignInHandler).Methods("POST")
	a.Router.HandleFunc("/api/sign-up", a.SignUpHandler).Methods("POST")
	a.Router.HandleFunc("/api/profile", a.ProfileHandler).Methods("GET")
	a.Router.HandleFunc("/api/transfer/autocomplete", a.AutoCompleteHandler).Methods("POST")
	a.Router.HandleFunc("/api/transfer/create", a.TransferHandler).Methods("POST")
	a.Router.HandleFunc("/api/transfer/history", a.HistoryHandler).Methods("POST")
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
