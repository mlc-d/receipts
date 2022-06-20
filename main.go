package main

import (
	"log"
	"net/http"
	"recibosV2/controllers"
	"recibosV2/data"
	"recibosV2/models"
)

func main() {
	data.Connect()
	data.Gdb.AutoMigrate(
		models.Person{},
	)
	http.HandleFunc("/users", controllers.PersonController)
	log.Println("Server up. Listening on port :1998")
	log.Fatal(http.ListenAndServe(":1998", nil))
}
