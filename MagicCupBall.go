package main

import "fmt"

func swap(arr *[3]int, move int) {
	var temp int
	switch move {
	case 1:
		temp = arr[0]
		arr[0] = arr[1]
		arr[1] = temp
	case 2:
		temp = arr[1]
		arr[1] = arr[2]
		arr[2] = temp
	case 3:
		temp = arr[0]
		arr[0] = arr[2]
		arr[2] = temp
	}
}

func main() {
	var amount int
	cup := [3]int{1, 0, 0}

	fmt.Scan(&amount)
	actions := make([]int, amount)

	for i := 0; i < amount; i++ {
		fmt.Scan(&actions[i])
	}

	for i := 0; i < amount; i++ {
		swap(&cup, actions[i])
	}

	fmt.Println(cup)

	if cup[0] == 1 {
		fmt.Println("ball is in cup 1")
	} else if cup[1] == 1 {
		fmt.Println("ball is in cup 2")
	} else {
		fmt.Println("ball is in cup 3")
	}
}
