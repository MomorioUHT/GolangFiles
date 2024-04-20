package main

import "fmt"

func main() {
	var fname, lname, a1, a2, gender, tel string
	fmt.Scanf("%s\n", &fname)
	fmt.Scanf("%s\n", &lname)
	fmt.Scanf("%s %s\n", &a1, &a2)
	fmt.Scanf("%s\n", &gender)
	fmt.Scanf("%s", &tel)

	fmt.Println("--- Customer Detail ---")
	fmt.Printf("Name : %s %s\n", fname, lname)
	fmt.Printf("Address : %s %s\n", a1, a2)
	fmt.Printf("Gender : %s\n", gender)
	fmt.Printf("Tel : %s\n", tel)
}
