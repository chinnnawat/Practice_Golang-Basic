package main

import "fmt"

func grad(value int){
	if value>=80 {
		fmt.Println("A")
	} else if value>=70 {
		fmt.Println("B")
	} else if value>=60 {
		fmt.Println("C")
	} else if value>=50 {
		fmt.Println("D")
	} else if value>=0 {
		fmt.Println("ติด F")
	} else {
		fmt.Println("อย่าเอ๋อ")
	}
}

func main() {
	var score int
	fmt.Printf("Enter your Score = ")
	fmt.Scanf("%d", &score)
	grad(score)
}