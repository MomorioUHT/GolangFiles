package main

import (
	"fmt"
)

func main() {
	var arr [5]float64
	var sum float64 = 0
	for i := 0; i < 5; i++ {
		fmt.Scan(&arr[i])
		sum += arr[i]
	}
	avg := sum / float64(len(arr))

	fmt.Printf("THAI = %.1f\n", arr[0])
	fmt.Printf("MATH = %.1f\n", arr[1])
	fmt.Printf("ENGLISH = %.1f\n", arr[2])
	fmt.Printf("SCIENCE = %.1f\n", arr[3])
	fmt.Printf("SPORT = %.1f\n", arr[4])
	fmt.Printf("---\n")
	fmt.Printf("GPA = %.1f", avg)

}
