package routers

import (
	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
	"framework/common"
	"framework/controllers"
)

// SetCountryRoutes configures routes for task entity
func SetCountryRoutes(router *mux.Router) *mux.Router {
	countryRouter := mux.NewRouter()
	countryRouter.HandleFunc("/country", controllers.CreateCountry).Methods("POST")
	countryRouter.HandleFunc("/country/{id}", controllers.UpdateCountry).Methods("PUT")
	countryRouter.HandleFunc("/country", controllers.GetCountries).Methods("GET")
	countryRouter.HandleFunc("/country/{id}", controllers.GetCountryByID).Methods("GET")
	countryRouter.HandleFunc("/country/users/{id}", controllers.GetCountriesByUser).Methods("GET")
	countryRouter.HandleFunc("/country/{id}", controllers.DeleteCountry).Methods("DELETE")
	router.PathPrefix("/country").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(countryRouter),
	))
	return router
}
