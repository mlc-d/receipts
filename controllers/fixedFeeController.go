package controllers

import (
	"encoding/json"
	"net/http"
	"recibosV2/errorh"
	"recibosV2/models"
)

func FixedFeeGetController(w http.ResponseWriter, r *http.Request) {
	var fixedFees []models.FixedFee
	var err error
	fixedFees = models.GetFixedFees()
	payload, err := json.Marshal(fixedFees)
	errorh.Handle(err)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
	return
}

func FixedFeePostController(w http.ResponseWriter, r *http.Request) {
	fixedFee := new(models.FixedFee)
	var err error
	err = json.NewDecoder(r.Body).Decode(&fixedFee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.CreateFixedFee(*fixedFee)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("201 - created"))
	return
}

func FixedFeePatchController(w http.ResponseWriter, r *http.Request) {
	fixedFee := new(models.FixedFee)
	var err error
	err = json.NewDecoder(r.Body).Decode(&fixedFee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.UpdateFixedFee(*fixedFee)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 - ok"))
	return
}

func FixedFeeDeleteController(w http.ResponseWriter, r *http.Request) {
	fixedFee := new(models.FixedFee)
	var err error
	err = json.NewDecoder(r.Body).Decode(&fixedFee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.DeleteFixedFee(*fixedFee)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 - ok"))
	return
}

func FixedFeeController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		FixedFeeGetController(w, r)
	case "POST":
		FixedFeePostController(w, r)
	case "PATCH":
		FixedFeePatchController(w, r)
	case "DELETE":
		FixedFeeDeleteController(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("405 - Method Not Allowed"))
	}
}
