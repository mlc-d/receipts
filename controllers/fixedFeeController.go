package controllers

import (
	"encoding/json"
	"net/http"
	"recibosV2/errorh"
	"recibosV2/models"
)

func FixedFeeController(w http.ResponseWriter, r *http.Request) {
	fixedFee := new(models.FixedFee)
	var fixedFees []models.FixedFee
	var err error
	switch r.Method {
	case "GET":
		fixedFees = models.GetFixedFees()
		payload, err := json.Marshal(fixedFees)
		errorh.Handle(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(payload)
		return
	case "POST":
		err = json.NewDecoder(r.Body).Decode(&fixedFee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		models.CreateFixedFee(*fixedFee)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("201 - created"))
		return
	case "PATCH":
		err = json.NewDecoder(r.Body).Decode(&fixedFee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		models.UpdateFixedFee(*fixedFee)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - ok"))
		return
	case "DELETE":
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
}
