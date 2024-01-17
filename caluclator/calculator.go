package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// go => เวลาเรียกใช้ func จาก library จะให้มา 2 ตัวคือ 1.สิ่งที่ฟังก์ชันนั้นทำ 2. error ที่เกิดขึ้น เสมอ

// การสร้าง bufio.Reader โดยกำหนดให้ os.Stdin เป็นแหล่งข้อมูล.
// os.Stdin คือตัวแปรของ Go ที่ใช้แทน Standard Input (stdin) หรือข้อมูลที่รับเข้ามาจากคีย์บอร์ด
var reader = bufio.NewReader(os.Stdin)

func getInput(promp string)float64 {
	// หากไม่แน่ใจว่า data type เป็น type อะไร ให้ใช้เป็น %v
	fmt.Printf("%v", promp)

	// _ => void ในภาษา c
	// _ => ใช้รับค่า error ที่เกิดขึ้น
	// ReadString('\n') คือการให้ reader อ่านข้อมูลจนกว่าจะพบตัว \n (newline)
	input, _ := reader.ReadString('\n')

	// strconv.ParseFloat แปลง str => float
	// TrimSpace ลบช่องว่างระหว่าง ขอบเขต string
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	// if err != nil (มี error)
	if err != nil {
		message,_ := fmt.Printf("%v number only", value)
		// panic ใช้เหมือน print แต่ใช้แสดง ข้อความที่ error
		panic(message)
	}
	return value
}

func getOperator() string{
	fmt.Printf("oparetor = (+,-,*,/) : ")
	op,_ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func addValue(value1 , value2 float64) float64{
	return value1+value2	
}

func minusValue(value1 , value2 float64)float64{
	return value1-value2
}

func multiValue(value1, value2 float64)float64{
	return value1*value2
}

func devideValue(value1, value2 float64)float64{
	return value1/value2
}

func main() {
	value1 := getInput("value1 = ")
	value2 := getInput("value2 = ")

	var result float64

	// switch expection ที่สร้างขึ้นมา; expection ที่เรียกใช้งาน {}
	switch oparetor := getOperator(); oparetor{
	case "+":
		result = addValue(value1, value2)
	case "-":
		result = minusValue(value1, value2)
	case "*":
		result = multiValue(value1, value2)
	case "/":
		result = devideValue(value1, value2)
	default :
		fmt.Println("invalid")
	}

	fmt.Println("result = ",result)

}