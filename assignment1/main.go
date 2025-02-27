package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range s {
		if checkOdd(num) {
			fmt.Println(num, "is odd")
		} else {
			fmt.Println(num, "is even")
		}
	}
}

func checkOdd(num int) bool {
	return num%2 != 0
}
