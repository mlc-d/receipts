package controllers

import (
	"encoding/json"
	"net/http"
	"recibosV2/errorh"
	"recibosV2/models"
)

func ServiceGetController(w http.ResponseWriter, r *http.Request) {
	var services []models.Service
	var err error
	services = models.GetServices()
	payload, err := json.Marshal(services)
	errorh.Handle(err)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
	return
}

func ServicePostController(w http.ResponseWriter, r *http.Request) {
	service := new(models.Service)
	var err error
	err = json.NewDecoder(r.Body).Decode(&service)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.CreateService(*service)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("201 - created"))
	return
}

func ServicePatchController(w http.ResponseWriter, r *http.Request) {
	service := new(models.Service)
	var err error
	err = json.NewDecoder(r.Body).Decode(&service)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.UpdateService(*service)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 - ok"))
	return
}

func ServiceDeleteController(w http.ResponseWriter, r *http.Request) {
	service := new(models.Service)
	var err error
	err = json.NewDecoder(r.Body).Decode(&service)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.DeleteService(*service)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 - ok"))
	return
}

func ServiceController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ServiceGetController(w, r)
	case "POST":
		ServicePostController(w, r)
	case "PATCH":
		ServicePatchController(w, r)
	case "DELETE":
		ServiceDeleteController(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("405 - Method Not Allowed"))
	}
}
