package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, its my hw0!")
	fmt.Println(Fizzbuzz(3))
	fmt.Println(IsPrime(3))
	fmt.Println(IsPalindrome("go"))
}

func Fizzbuzz(n int) string {

	if (n % 3 == 0 && n % 5 == 0) {
		return "FizzBuzz"
	} else if (n % 3 == 0) {
		return "Fizz"
	} else if (n % 5 == 0) {
		return "Buzz"
	} else {
		return ""
	}

}

func IsPrime(n int) bool {
	if n < 2 {
		return false
	} else if 2 == n || 3 == n {
		return true
	} else {
		limit := math.Sqrt(float64(n))
		for i := 2; i <= int(limit); i++ {
			if n % i == 0 {
				return false
			}
		}
		return true
	}
}

func IsPalindrome(s string) bool {
	reverse := ""

	for i := len(s) - 1; i >= 0; i-- {
	  reverse += string(s[i])
	}
	  
	for i := 0; i < len(s); i++ {
		if s[i] != reverse[i] {
			return false
		}
	}
	
	return true
}