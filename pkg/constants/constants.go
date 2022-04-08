package constants

import "log"

var choice = map[int]string{
	1: "Primes",
	2: "None",
}

func ResolveKey(key int) string {
	value, found_or_not := choice[key]

	if !found_or_not {
		log.Fatalf("Error in Package Constants: Key Invalid")
	}
	return value
}
