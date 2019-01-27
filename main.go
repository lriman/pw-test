package main

import (
	"log"
)

func main() {
	a := new(App)

	a.Initialize("127.0.0.1", "myuser", "mypass", "mydb")
	log.Println("Database success connection")

	log.Println("Service is open on 8001 port")
	a.Run(":8001")
}

