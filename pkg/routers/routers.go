package routers

import (
	"sample_project/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("api/prime_list/{upper_limit}", controllers.PrimeList).Methods("GET")
	router.HandleFunc("api/find_factors/{num}", controllers.FindFactors).Methods("GET")
	router.HandleFunc("api/int_to_binary/{num}", controllers.IntToBinary).Methods("GET")
	router.HandleFunc("api/random_binary/{bits}", controllers.RandomBinary).Methods("GET")
}
