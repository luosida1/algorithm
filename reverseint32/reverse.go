package main

import "fmt"

func reverse(x int) int {
	re := 0
	for i := 0; x != 0; i++ {
		pop := x % 10
		x /= 10

		re = re*10 + pop

		if re > 0x7fffffff || re < -0x80000000 {
			return 0
		}
		fmt.Printf("re = %d, x = %d\n", re, x)
	}

	return re
}

func main() {
	//8463847412
	inputNum := -2147483649
	reverseNum := reverse(inputNum)
	fmt.Printf("origin = %d, reverse = %d\n", inputNum, reverseNum)
}
