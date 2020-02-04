package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	re := 0
	for x > re {
		re = x%10 + re*10
		x /= 10
	}

	return x == re || x == re/10
}

func main() {

	result := isPalindrome(1232111)
	fmt.Println(result)

}
