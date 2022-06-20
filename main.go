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
		models.Service{},
		models.Receipt{},
		models.Obligation{},
		models.FixedFee{},
		models.ExtraFee{},
		models.Payment{},
	)
	http.HandleFunc("/users", controllers.PersonController)
	http.HandleFunc("/services", controllers.ServiceController)
	http.HandleFunc("/fixedfees", controllers.FixedFeeController)
	http.HandleFunc("/extrafees", controllers.ExtraFeeController)
	http.HandleFunc("/receipts", controllers.ReceiptController)
	http.HandleFunc("/obligations", controllers.ObligationController)
	http.HandleFunc("/payments", controllers.PaymentController)

	log.Println("Server up. Listening on port :1998")
	log.Fatal(http.ListenAndServe(":1998", nil))
}
