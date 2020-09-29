// SOLUTION TO QUESTION 2B

package main

import (
	"fmt"
	"math/big"
)

// Function to Check for the Power of two
func powerOfTwo(n int64) bool {
	result := n != 0 && (n&(n-1)) == 0
	return result
}

// Function to Check if p is a proth number
// for k, an odd integer and < 2^n
func prothNumber(p int64) bool {

	var k int64 = 1

	for k < (p / k) {
		if p%k == 0 {
			if powerOfTwo(p / k) {
				return true
			}
		}
		k += 2
	}
	return false
}

// Utility Function that does modular
// integer arithmetic for large Numbers
func powInt(x, y, z int64) int64 {

	var (
		i = big.NewInt(x)
		e = big.NewInt(y)
		k = big.NewInt(z)
	)

	// i ^ e modulus k
	i.Exp(i, e, k)

	// Convert from Big Int to int.64
	return i.Int64()
}

// Check if given number is a proth prime
func prothPrime(p int64) bool {

	var (
		pm1 int64 = p - 1
		x   int64 = pm1 / 2
		a   int64 = 1
	)

	for a < p {

		// Modular Exponentiation for proth primality test
		// where, a ^ ((p - 1) / 2) equiv 1 (mod p)
		if powInt(a, x, p) == pm1 {
			return true
		}
		a++
	}
	return false
}

// Function to Check if input
// is Proth number and Proth Prime
func main() {

	fmt.Print("Enter a number: ")
	var p int64
	fmt.Scanln(&p)

	if prothNumber(p - 1) {

		if prothPrime(p) == true {

			fmt.Println("Proth Number and Proth Prime")

		} else {

			fmt.Println("Proth Number but not a Proth Prime")

		}
	} else {
		fmt.Println("Not a Proth Number")
	}

}
