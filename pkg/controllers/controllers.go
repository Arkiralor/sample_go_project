package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sample_project/pkg/binaries"
	"sample_project/pkg/factors"
	"sample_project/pkg/prime_numbers"
	"strconv"

	"github.com/gorilla/mux"
)

func PrimeList(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	upper_limit_str := params["upper_limit"]
	upper_limit, err := strconv.Atoi(upper_limit_str)
	if err != nil {
		log.Printf("Error: %v", err.Error())
		panic(err)
	}

	prime_list := prime_numbers.FindListOfPrimes(upper_limit)

	var resp map[string]interface{} = map[string]interface{}{
		"function":      "Find the list of all Prime Numbers <= 'n'.",
		"input":         upper_limit,
		"prime_numbers": prime_list,
		"length":        len(prime_list),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func FindFactors(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	num_str := params["num"]
	num, err := strconv.Atoi(num_str)
	if err != nil {
		log.Printf("Error: %v", err.Error())
		panic(err)
	}
	factor_list := factors.FindFactors(num)
	var resp map[string]interface{} = map[string]interface{}{
		"function":      "Find the list of all Factors of 'n'.",
		"input":         num,
		"prime_numbers": factor_list,
		"length":        len(factor_list),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func IntToBinary(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	num_str := params["num"]
	num, err := strconv.Atoi(num_str)
	if err != nil {
		log.Printf("Error: %v", err.Error())
		panic(err)
	}
	binary_number := binaries.Int64ToBinary(num)
	var resp map[string]interface{} = map[string]interface{}{
		"function":      "Convert 'n' in decimal to 'n' in binary.",
		"input":         num,
		"prime_numbers": binary_number,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func RandomBinary(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bits_str := params["bits"]
	bits, err := strconv.Atoi(bits_str)
	if err != nil {
		log.Printf("Error: %v", err.Error())
		panic(err)
	}
	binary_number := binaries.GenerateRandomBinaryNumber(bits)
	var resp map[string]interface{} = map[string]interface{}{
		"function":      "Generate a random binary number of 'size' bits.",
		"input":         bits,
		"prime_numbers": binary_number,
		"length":        bits,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
