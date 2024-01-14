package main

import "fmt"

func numCount(){
	count := 10
	count2 := 20
	for i := 1; i <= count; i++ {
		fmt.Println("i =",i)
	}

	for j := 11; j <= count2; j++ {
		fmt.Println("j =", j)
	}
}

func stringCount(){
	var input string
	for {
		fmt.Printf("Enter your message = ")
		fmt.Scanf("%s", &input)
		fmt.Println("Output =", input)
		if input == "exit" {
			break
		}
	}

}

func main() {
	// numCount()
	stringCount()
}