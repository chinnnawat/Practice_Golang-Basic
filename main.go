package main

import "fmt"

var numberInt1, numberInt2 int = 2000, 1000;
var msg string = "Hello"

func main() {

	// โชว์ชนิดตัวแปร
	// := เมื่อใช้ระบุใน int จะเป็นแบบ float64
	numberfloat := 25.4
	fmt.Println(numberInt1)
	fmt.Println(numberInt2)
	fmt.Println(numberfloat)
	fmt.Println(msg)

	// การรวมกันของตัวแปร
	// 1. การบวกของ int+int = int
	fmt.Println(numberInt1+numberInt2)
	// 2. การบวกของ int + float(float64) = float(float64)
	fmt.Println(float64(numberInt1)+(numberfloat))
	// 3. การบวกกันของ string
	fmt.Println(msg+"Chin")
	// 4. การต่อกันของข้อความ
	fmt.Println("My money =",numberInt1)
}