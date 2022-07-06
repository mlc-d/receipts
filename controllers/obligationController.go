package controllers

import (
	"encoding/json"
	"net/http"
	"recibosV2/errorh"
	"recibosV2/models"
)

func ObligationGetController(w http.ResponseWriter, r *http.Request) {
	var obligations []models.Obligation
	var err error
	obligations = models.GetObligations()
	payload, err := json.Marshal(obligations)
	errorh.Handle(err)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
	return
}

func ObligationPostController(w http.ResponseWriter, r *http.Request) {
	obligation := new(models.Obligation)
	var err error
	err = json.NewDecoder(r.Body).Decode(&obligation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.CreateObligation(*obligation)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("201 - created"))
	return
}

func ObligationPatchController(w http.ResponseWriter, r *http.Request) {
	obligation := new(models.Obligation)
	var err error
	err = json.NewDecoder(r.Body).Decode(&obligation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.UpdateObligation(*obligation)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 - ok"))
	return
}

func ObligationDeleteController(w http.ResponseWriter, r *http.Request) {
	obligation := new(models.Obligation)
	var err error
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

func ObligationController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ObligationGetController(w, r)
	case "POST":
		ObligationPostController(w, r)
	case "PATCH":
		ObligationPatchController(w, r)
	case "DELETE":
		ObligationDeleteController(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("405 - Method Not Allowed"))
	}
}
