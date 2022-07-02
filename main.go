package main

import (
	"log"
	"net/http"
	"recibosV2/data"
	"recibosV2/models"
	"recibosV2/routes"
)

func main() {
	data.Connect()
	data.Gdb.AutoMigrate(
		models.Person{},
		models.Service{},
		models.Receipt{},
		models.Obligation{},
		models.FixedFee{},
		models.ExtraFee{},
		models.Payment{},
	)

	routes.ServeRoutes()

	log.Println("Server up. Listening on port :1998")
	log.Fatal(http.ListenAndServe(":1998", nil))
}
