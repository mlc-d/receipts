package controllers

import (
	"encoding/json"
	"net/http"
	"recibosV2/errorh"
	"recibosV2/models"
)

func ServiceController(w http.ResponseWriter, r *http.Request) {
	service := new(models.Service)
	var services []models.Service
	var err error
	switch r.Method {
	case "GET":
		services = models.GetServices()
		payload, err := json.Marshal(services)
		errorh.Handle(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(payload)
		return
	case "POST":
		err = json.NewDecoder(r.Body).Decode(&service)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		models.CreateService(*service)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("201 - created"))
		return
	case "PATCH":
		err = json.NewDecoder(r.Body).Decode(&service)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		models.UpdateService(*service)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - ok"))
		return
	case "DELETE":
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
}
