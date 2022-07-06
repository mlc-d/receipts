package controllers

import (
	"encoding/json"
	"net/http"
	"recibosV2/errorh"
	"recibosV2/models"
)

func ExtraFeeGetController(w http.ResponseWriter, r *http.Request) {
	var extraFees []models.ExtraFee
	var err error
	extraFees = models.GetExtraFees()
	payload, err := json.Marshal(extraFees)
	errorh.Handle(err)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
	return
}

func ExtraFeePostController(w http.ResponseWriter, r *http.Request) {
	extraFee := new(models.ExtraFee)
	var err error
	err = json.NewDecoder(r.Body).Decode(&extraFee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.CreateExtraFee(*extraFee)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("201 - created"))
	return
}

func ExtraFeePatchController(w http.ResponseWriter, r *http.Request) {
	extraFee := new(models.ExtraFee)
	var err error
	err = json.NewDecoder(r.Body).Decode(&extraFee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.UpdateExtraFee(*extraFee)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 - ok"))
	return
}

func ExtraFeeDeleteController(w http.ResponseWriter, r *http.Request) {
	extraFee := new(models.ExtraFee)
	var err error
	err = json.NewDecoder(r.Body).Decode(&extraFee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.DeleteExtraFee(*extraFee)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 - ok"))
	return
}

func ExtraFeeController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ExtraFeeGetController(w, r)
	case "POST":
		ExtraFeePostController(w, r)
	case "PATCH":
		ExtraFeePatchController(w, r)
	case "DELETE":
		ExtraFeeDeleteController(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("405 - Method Not Allowed"))
		return
	}
}
