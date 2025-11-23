package main

import (
	"errors"
	"fmt"
)

func sumOfNumbers(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := sumOfNumbers(numbers)
	const expected int = 45
	if result != expected {
		testError := errors.New("sum of numbers failed")
		fmt.Println("Test Failed: ", testError)
	}
}
