package main

import "fmt"

// jwtToken:="qwertyuiop" can't use := outside main
var jwtToken string = "qwertyuiop" // is valid

const api_key = "asdfghjkl123456" // Public can be access anywhere global

func main() {

	var name string = "Nayan"
	fmt.Println(name)
	fmt.Printf("Variable Type is %T \n", name)

	var isGood bool = true
	fmt.Println(isGood)
	fmt.Printf("Variable Type is %T \n", isGood)

	var uintVal uint8 = 255
	fmt.Println(uintVal)
	fmt.Printf("Variable Type is %T \n", uintVal)

	var floatVal float64 = 255.56784353242317811323
	fmt.Println(floatVal)
	fmt.Printf("Variable Type is %T \n", floatVal)

	//Default Values
	var intDefault int
	var stringDefault string
	var uintDefault uint
	var boolDefault bool
	fmt.Printf("Default Int Value : %v\n", intDefault)
	fmt.Printf("Default String Value : %v\n", stringDefault)
	fmt.Printf("Default Uint Value : %v\n", uintDefault)
	fmt.Printf("Default Bool Value : %v\n", boolDefault)

	//Impicit Type
	var userName = "iamnayan31"
	fmt.Println(userName)

	// No var style
	numberOfYear := 5
	fmt.Println(numberOfYear)

	//Using Constant
	fmt.Println(api_key)
	fmt.Printf("Variable Type is : %T", api_key)

}
