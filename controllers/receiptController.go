package controllers

import (
	"encoding/json"
	"net/http"
	"recibosV2/errorh"
	"recibosV2/models"
)

func ReceiptGetController(w http.ResponseWriter, r *http.Request) {
	var receipts []models.Receipt
	var err error
	receipts = models.GetReceipts()
	payload, err := json.Marshal(receipts)
	errorh.Handle(err)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
	return
}

func ReceiptPostController(w http.ResponseWriter, r *http.Request) {
	receipt := new(models.Receipt)
	var err error
	err = json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.CreateReceipt(*receipt)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("201 - created"))
	return
}

func ReceiptPatchController(w http.ResponseWriter, r *http.Request) {
	receipt := new(models.Receipt)
	var err error
	err = json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.UpdateReceipt(*receipt)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 - ok"))
	return
}

func ReceiptDeleteController(w http.ResponseWriter, r *http.Request) {
	receipt := new(models.Receipt)
	var err error
	err = json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.DeleteReceipt(*receipt)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 - ok"))
	return
}

func ReceiptController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ReceiptGetController(w, r)
	case "POST":
		ReceiptPostController(w, r)
	case "PATCH":
		ReceiptPatchController(w, r)
	case "DELETE":
		ReceiptDeleteController(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("405 - Method Not Allowed"))
	}
}
