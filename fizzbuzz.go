package fizzbuzz

import "fmt"

func Fizzbuzz(n int) string {
	if n%15 == 0 {
		return "FizzBuzz"
	}
	if n%5 == 0 {
		return "Buzz"
	}
	if n%3 == 0 {
		return "Fizz"
	}
	return fmt.Sprint(n)
}

func FizzBuzzInterval(start int, end int) string {
	var s string
	for start <= end {
		s += Fizzbuzz(start)
		start++
	}
	return s
}
