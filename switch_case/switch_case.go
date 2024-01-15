package main

import (
	"fmt"
)

// func switchCase(input int) {
// 	switch input{
// 	case 1:
// 		fmt.Println("one")
// 	case 2:
// 		fmt.Println("two")
// 	case 3:
// 		fmt.Println("three")
// 	default :
// 		fmt.Println("invalid")
// 	}
// }

func switchColor(input string){
	switch input{
	case "blue":
		fmt.Println("blue = #0000FF")
	case "green":
		fmt.Println("green = #008000")
	case "pink" :
		fmt.Println("pink = #FFC0CB")
	case "yellow" :
		fmt.Println("yellow = #FFFF00")
	default :
		fmt.Println("invalid")
	}
}

func main() {
	// number
	var input string
	// fmt.Printf("Enter your Number(1-3) =")
	// fmt.Scanf("%d", &input)
	// switchCase(input)

	// color
	fmt.Printf("Enter your Color = ")
	fmt.Scanf("%s", &input)
	switchColor(input)
}