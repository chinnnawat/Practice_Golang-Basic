package main

import "fmt"

// productName มีขนาด 4 เป็นลักษณะ string
// [] = กำหนดขนาด (basket)
var productName[4]string
// var price = [4]float32{1000,200,3000,4000}

func main() {
	productName[0] = "apple"
	productName[1] = "banana"
	productName[2] = "water melon"
	productName[3] = "cake"

	price := [4]float32{1000,200,3000,4000}

	fmt.Println(productName)
	fmt.Println(price)
}