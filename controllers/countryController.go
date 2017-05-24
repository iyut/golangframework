package controllers

import (
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"framework/common"
	"framework/data"
)

// CreateTask insert a new Task document
// Handler for HTTP Post - "/country
func CreateCountry(w http.ResponseWriter, r *http.Request) {
	var dataResource CountryResource

	// Decode the incoming Country json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Country data",
			500,
		)
		return
	}
	country := &dataResource.Data

	val := common.GetContextAuth(r)
	if val != "" {
		country.CreatedBy = val
	}else{
		country.CreatedBy = "noone"
	}
	repo := &data.CountryRepository{C: "MS_Country"}
	// Insert a country document
	repo.Create(country)
	j, err := json.Marshal(CountryResource{Data: *country})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)

}

// GetTasks returns all Task document
// Handler for HTTP Get - "/countries"
func GetCountries(w http.ResponseWriter, r *http.Request) {
fmt.Println("test ok")
	repo := &data.CountryRepository{C: "MS_Country"}
	countries := repo.GetAll()
	j, err := json.Marshal(CountriesResource{Data: countries})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// GetTaskByID returns a single Task document by id
// Handler for HTTP Get - "/country/{id}"
func GetCountryByID(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	repo := &data.CountryRepository{C: "MS_Country"}
	country, err := repo.GetById(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)

		} else {
			common.DisplayAppError(
				w,
				err,
				"An unexpected error has occurred",
				500,
			)

		}
		return
	}
	j, err := json.Marshal(country)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// GetCountriesByUser returns all Countries created by a User
// Handler for HTTP Get - "/countries/users/{id}"
func GetCountriesByUser(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	user := vars["id"]

	repo := &data.CountryRepository{C: "MS_Country"}
	countries := repo.GetByUser(user)
	j, err := json.Marshal(CountriesResource{Data: countries})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// UpdateTask update an existing Country  document
// Handler for HTTP Put - "/country/{id}"
func UpdateCountry(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	var dataResource CountryResource
	// Decode the incoming Task json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Task data",
			500,
		)
		return
	}
	country := &dataResource.Data
	country.Id = id

	repo := &data.CountryRepository{C: "MS_Country"}
	// Update an existing Country document
	if err := repo.Update(country); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

// DeleteTask deelete an existing Country document
// Handler for HTTP Delete - "/country/{id}"
func DeleteCountry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	repo := &data.CountryRepository{C: "MS_Country"}
	// Delete an existing Task document
	err := repo.Delete(id)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
