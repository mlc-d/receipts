package controllers

import (
	"encoding/json"
	"net/http"
	"recibosV2/data"
	"recibosV2/errorh"
	"recibosV2/models"
	"recibosV2/utils"
)

func PaymentGetController(w http.ResponseWriter, r *http.Request) {
	var payments []models.Payment
	var err error
	payments = models.GetPayments()
	payload, err := json.Marshal(payments)
	errorh.Handle(err)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(payload)
	return
}

func PaymentPostController(w http.ResponseWriter, r *http.Request) {
	var err error
	payment := new(models.Payment)
	var payments []models.Payment
	err = json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	receipt := models.GetReceipt(models.Receipt{Id: payment.ReceiptID})
	payments = calculatePayment(receipt)
	for _, v := range payments {
		models.CreatePayment(v)
	}
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("201 - created"))
	return
}

func PaymentPatchController(w http.ResponseWriter, r *http.Request) {
	payment := new(models.Payment)
	var err error
	err = json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.UpdatePayment(*payment)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("200 - ok"))
	return
}

func PaymentDeleteController(w http.ResponseWriter, r *http.Request) {
	var err error
	payment := new(models.Payment)
	err = json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.DeletePayment(*payment)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("200 - ok"))
	return
}

func PaymentController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		PaymentGetController(w, r)
	case "POST":
		PaymentPostController(w, r)
	case "PATCH":
		PaymentPatchController(w, r)
	case "DELETE":
		PaymentDeleteController(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("405 - Method Not Allowed"))
	}
}

func calculatePayment(receipt models.Receipt) (payments []models.Payment) {
	// connect database
	db := data.Gdb

	// variables to store the amount of fixed and extra fees
	fixedFees := make([]map[string]interface{}, 0)
	extraFees := make([]map[string]interface{}, 0)

	// retrieve the values of fees
	db.Model(&models.FixedFee{}).Where("service_id = ?", receipt.ServiceID).Select("id, person_id, amount").Find(&fixedFees)
	db.Model(&models.ExtraFee{}).Where("service_id = ?", receipt.ServiceID).Select("id, person_id, amount").Find(&extraFees)

	// get the combined amount
	var totalFees float32
	feesCreditors := make(map[string]interface{})

	for _, v := range fixedFees {
		feesCreditors[utils.StringValue(v["person_id"])] = v["amount"]
		totalFees += utils.Float32Value(v["amount"])
	}
	for _, v := range extraFees {
		feesCreditors[utils.StringValue(v["person_id"])] = v["amount"]
		totalFees += utils.Float32Value(v["amount"])
	}

	// get person_id of persons who have to pay for the service
	creditors := make([]map[string]interface{}, 0)
	db.Model(&models.Obligation{}).Where("service_id = ?", receipt.ServiceID).Select("person_id").Scan(&creditors)

	// create the array of payments
	if len(creditors) > 0 {
		amount := (receipt.Amount - totalFees) / float32(len(creditors))
		for _, v := range creditors {
			if feesCreditors[utils.StringValue(v["person_id"])] == nil {
				feesCreditors[utils.StringValue(v["person_id"])] = 0.0
			}
			payments = append(payments, models.Payment{
				ReceiptID: receipt.Id,
				PersonID:  utils.Uint8Value(v["person_id"]),
				Amount:    amount + utils.Float32Value(feesCreditors[utils.StringValue(v["person_id"])]),
			})
		}
	}

	// if there are any fixed fees, append them
	if len(fixedFees) > 0 {
		for _, v := range fixedFees {
			payments = append(payments, models.Payment{
				ReceiptID: receipt.Id,
				PersonID:  utils.Uint8Value(v["person_id"]),
				Amount:    utils.Float32Value(v["amount"]),
			})
		}
	}
	return
}
