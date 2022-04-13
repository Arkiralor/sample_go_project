package routers

import (
	"sample_project/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("api/v1/prime_list/{upper_limit}", controllers.PrimeList).Methods("GET")
	router.HandleFunc("api/v1/find_factors/{num}", controllers.FindFactors).Methods("GET")
	router.HandleFunc("api/v1/int_to_binary/{num}", controllers.IntToBinary).Methods("GET")
	router.HandleFunc("api/v1/random_binary/{bits}", controllers.RandomBinary).Methods("GET")
}
