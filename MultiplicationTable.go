package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	for i := 1; i < 13; i++ {
		fmt.Printf("%d x %d = %d\n", number, i, number*i)
	}
}
