package main

import "fmt"

func main() {

	var fruitlist [4]string

	fruitlist[0] = "Apple"
	fruitlist[1] = "Tomato"
	fruitlist[3] = "Peach"

	fmt.Println("Fruit list is:", fruitlist)
	fmt.Println("Fruit list length is:", len(fruitlist))

	var veglist = [5]string{"potato", "Beans", "mushroom"}
	fmt.Println("vegy list is:", veglist)
	fmt.Println("vegy list length is:", len(veglist))
}
