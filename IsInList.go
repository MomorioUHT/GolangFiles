package main

import "fmt"

func isInList(ls []int, x int) bool {
	for _, value := range ls {
		if value == x {
			return true
		}
	}
	return false
}

func main() {
	var amount, find int
	fmt.Scan(&amount)

	arr := make([]int, amount)

	for i := 0; i < amount; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Scan(&find)

	if isInList(arr, find) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
