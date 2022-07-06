package controllers

import (
	"encoding/json"
	"net/http"
	"recibosV2/errorh"
	"recibosV2/models"
)

func PersonGetController(w http.ResponseWriter, r *http.Request) {
	var persons []models.Person
	var err error
	persons = models.GetPersons()
	payload, err := json.Marshal(persons)
	errorh.Handle(err)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
	return
}

func PersonPostController(w http.ResponseWriter, r *http.Request) {
	person := new(models.Person)
	var err error

	err = json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.CreatePerson(*person)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("201 - created"))
	return
}

func PersonPatchController(w http.ResponseWriter, r *http.Request) {
	person := new(models.Person)
	var err error
	err = json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.UpdatePerson(*person)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 - ok"))
	return
}

func PersonDeleteController(w http.ResponseWriter, r *http.Request) {
	person := new(models.Person)
	var err error

	err = json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.DeletePerson(*person)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 - ok"))
	return
}

func PersonController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		PersonGetController(w, r)
	case "POST":
		PersonPostController(w, r)
	case "PATCH":
		PersonPatchController(w, r)
	case "DELETE":
		PersonDeleteController(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("405 - Method Not Allowed"))
	}
}
