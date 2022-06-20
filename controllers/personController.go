package controllers

import (
	"encoding/json"
	"net/http"
	"recibosV2/errorh"
	"recibosV2/models"
)

func PersonController(w http.ResponseWriter, r *http.Request) {
	person := new(models.Person)
	var persons []models.Person
	var err error
	switch r.Method {
	case "GET":
		persons = models.GetPersons()
		payload, err := json.Marshal(persons)
		errorh.Handle(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(payload)
	case "POST":
		err = json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		models.CreatePerson(*person)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("201 - created"))
	case "PATCH":
		err = json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		models.UpdatePerson(*person)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - ok"))
	}
}
