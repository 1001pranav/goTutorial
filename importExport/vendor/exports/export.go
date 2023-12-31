package export

import  (
	"math"
)

/*
	All the variables and functions which are exported from go lang needs to have Pascal Case
	Which means 1st letter needs to be capitalized
*/

//exported Functions
func PrimeNumbers(n int) []int {
	var primes []int;
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primes = append(primes, i);
		}
	}
	return primes;
}

// NonExported Functions
func isPrime(n int) bool {
	if n <  2 {
		return false;
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n % i == 0 {
			return false;
		}
	}

	return true;

}