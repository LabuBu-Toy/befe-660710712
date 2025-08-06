package main

import "fmt"

// email := "udompol_t@silpakorn.edu"

func main() {
	// var name string = "thanawit"
	var age int = 24

	email := "udompol_t@silpakorn.edu"
	gpa := 2.99

	firstname, lastname := "thanawit", "udompol"

	fmt.Printf("Name %s %s, age %d, email %s, gpa %.2f\n",
		firstname, lastname,
		age, email, gpa)
}
