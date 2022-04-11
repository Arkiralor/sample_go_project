package prime_numbers

import "fmt"

func FindListOfPrimes(upper_limit int) []int {
	var prime_list []int

	prime_list = append(prime_list, 2)

	for i := 3; i <= upper_limit; i += 1 {
		if CheckIfPrime(i) {
			fmt.Println("Prime number found: ", i)
			prime_list = append(prime_list, i)
		}
	}
	fmt.Print("\033[H\033[2J") // Clear the screen on Unix
	return prime_list
}

func CheckIfPrime(num int) bool {

	var is_prime, flag_var bool
	flag_var = false
	upper_limit := num / 2

	for i := 2; i <= upper_limit; i += 1 {
		if num%i == 0 {
			flag_var = true
		}
	}

	if flag_var {
		is_prime = false
	} else if !flag_var {
		is_prime = true
	}
	return is_prime
}
