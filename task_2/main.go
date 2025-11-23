package main

import (
	"errors"
	"fmt"
	"maps"
	"unicode"
)

func wordFrequencyCount(s string) map[rune]int {
    var counter map[rune]int = make(map[rune]int)
    for _, c := range s {
        if unicode.IsPunct(c) {
            continue
        }
        counter[unicode.ToLower(c)] += 1
    }
    return counter
}

func isPalindrome(s string) bool {
    left, right := 0, len(s) - 1
    var left_val , right_val rune
    for left < right {
        left_val, right_val = rune(s[left]), rune(s[right])
        if unicode.IsSpace(left_val) || unicode.IsPunct(left_val){
            left++
            continue
        }

        if unicode.IsSpace(right_val) || unicode.IsPunct(right_val){
            right--
            continue
        }

        if unicode.ToLower(left_val) != unicode.ToLower(right_val) {
            return false
        }

        left++
        right--
    }
    return true
}

func main()  {
    const s_counter string = "golang"
    var counter map[rune]int = wordFrequencyCount(s_counter)
    var expectedCounter map[rune]int = map[rune]int{'g': 2, 'o': 1, 'l': 1, 'a': 1, 'n': 1}

    if !maps.Equal(expectedCounter, counter){
        testError := errors.New("Counter does not match")
        fmt.Println("Test Failed: ", testError)
        return
    }

    const s_palindrome string = " r..ac  ec.,ar"
    var result bool = isPalindrome(s_palindrome)
    const expected bool = true
    if result != expected {
        testError := errors.New("isPalindrome failed")
        fmt.Println("Test Failed: ", testError)
        return
    }

    fmt.Println("All Tests Passed!");
}
