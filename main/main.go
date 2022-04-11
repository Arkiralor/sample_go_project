package main

import (
	"fmt"
	"os"
	"sample_project/pkg/console_input"
	"sample_project/pkg/constants"
	"sample_project/pkg/prime_numbers"
)

func main() {
	fmt.Printf("\nWhat is it that you wish to do today?\n1. Find Prime Numbers\n2. Do Nothing\n\n")
	inp := console_input.NumberFromCon()
	choice := constants.ResolveKey(inp)
	execute_choice(choice)
}

func execute_choice(choice string) {
	if choice == "Primes" {
		locate_primes()
	} else if choice == "None" {
		os.Exit(1)
	}
}

func locate_primes() {
	fmt.Printf("\nTill what number do you want to check for primes?\n")
	inp := console_input.NumberFromCon()
	list_of_primes := prime_numbers.FindListOfPrimes(inp)

	fmt.Printf("The list of Prime Numbers until %d is:\n%v\n", inp, list_of_primes)
}
