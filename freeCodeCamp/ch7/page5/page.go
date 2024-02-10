package page5

import (
	"fmt"
	"math"
)

func PrintPrimes(max int) {
	for n := 2; n < max+1; n++ {
		isPrime := true

		if n == 2 {
			fmt.Printf("%v is prime\n", n)
			isPrime = true
			continue
		}
		if n%2 == 0 {
			fmt.Printf("%v is not prime\n", n)
			isPrime = false
			continue
		}
		for i := 3; i <= int(math.Sqrt(float64(n)+1)); i++ {
			if n%i == 0 {
				fmt.Printf("%v is not prime\n", n)
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%v is prime\n", n)

		}
	}

}
