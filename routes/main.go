package routes

import (
	"net/http"
	"recibosV2/controllers"
)

func ServeRoutes() {
	http.HandleFunc("/users", controllers.PersonController)
	http.HandleFunc("/services", controllers.ServiceController)
	http.HandleFunc("/fixedfees", controllers.FixedFeeController)
	http.HandleFunc("/extrafees", controllers.ExtraFeeController)
	http.HandleFunc("/receipts", controllers.ReceiptController)
	http.HandleFunc("/obligations", controllers.ObligationController)
	http.HandleFunc("/payments", controllers.PaymentController)
}
