package controllers

import (
	"encoding/json"
	"net/http"
	"recibosV2/errorh"
	"recibosV2/models"
)

func ReceiptController(w http.ResponseWriter, r *http.Request) {
	receipt := new(models.Receipt)
	var receipts []models.Receipt
	var err error
	switch r.Method {
	case "GET":
		receipts = models.GetReceipts()
		payload, err := json.Marshal(receipts)
		errorh.Handle(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(payload)
	case "POST":
		err = json.NewDecoder(r.Body).Decode(&receipt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		models.CreateReceipt(*receipt)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("201 - created"))
	case "PATCH":
		err = json.NewDecoder(r.Body).Decode(&receipt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		models.UpdateReceipt(*receipt)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - ok"))
	case "DELETE":
		err = json.NewDecoder(r.Body).Decode(&receipt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		models.DeleteReceipt(*receipt)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - ok"))
	}
}
