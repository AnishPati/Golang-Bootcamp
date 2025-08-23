package main

import "fmt"

const LoginToken string = "AP"

func main() {
	var username string = "Anish"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type : %T \n", smallVal)

	var smallValue uint = 256
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type : %T \n", smallValue)

	var smallFloat float32 = 256.5645212155455
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)

	var smalFloat float64 = 256.5645212155455
	fmt.Println(smalFloat)
	fmt.Printf("Variable is of type : %T \n", smalFloat)

	//default values and some aliases
	var another int
	fmt.Println(another)
	fmt.Printf("Variable is of type : %T \n", another)

	//implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)

	//no var style
	numberofuser := 30000.0
	fmt.Println(numberofuser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n", LoginToken)

}
