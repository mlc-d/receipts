package controllers

import (
	"encoding/json"
	"net/http"
	"recibosV2/errorh"
	"recibosV2/models"
)

func ObligationController(w http.ResponseWriter, r *http.Request) {
	obligation := new(models.Obligation)
	var obligations []models.Obligation
	var err error
	switch r.Method {
	case "GET":
		obligations = models.GetObligations()
		payload, err := json.Marshal(obligations)
		errorh.Handle(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(payload)
		return
	case "POST":
		err = json.NewDecoder(r.Body).Decode(&obligation)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		models.CreateObligation(*obligation)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("201 - created"))
		return
	case "PATCH":
		err = json.NewDecoder(r.Body).Decode(&obligation)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		models.UpdateObligation(*obligation)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - ok"))
		return
	case "DELETE":
		err = json.NewDecoder(r.Body).Decode(&obligation)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		models.DeleteObligation(*obligation)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - ok"))
		return
	}
}
