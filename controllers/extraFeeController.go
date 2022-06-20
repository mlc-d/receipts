package controllers

import (
	"encoding/json"
	"net/http"
	"recibosV2/errorh"
	"recibosV2/models"
)

func ExtraFeeController(w http.ResponseWriter, r *http.Request) {
	extraFee := new(models.ExtraFee)
	var extraFees []models.ExtraFee
	var err error
	switch r.Method {
	case "GET":
		extraFees = models.GetExtraFees()
		payload, err := json.Marshal(extraFees)
		errorh.Handle(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(payload)
	case "POST":
		err = json.NewDecoder(r.Body).Decode(&extraFee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		models.CreateExtraFee(*extraFee)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("201 - created"))
	case "PATCH":
		err = json.NewDecoder(r.Body).Decode(&extraFee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		models.UpdateExtraFee(*extraFee)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - ok"))
	case "DELETE":
		err = json.NewDecoder(r.Body).Decode(&extraFee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		models.DeleteExtraFee(*extraFee)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - ok"))
	}
}
