package main

import "fmt"

func hello() {
	fmt.Println("Hello World")
}

// func รับค่า
// (ตัวแปร ชนิดตัวแปร,)
func plus(value1 int, value2 int){
	// var total int = value1+ value2
	total := value1 + value2
	fmt.Println("total=", total)
}

// func รับ และ return ค่า
// ชื่อfunc(ตัวแปร ชนิดตัวแปร, ตัวแปร ชนิดตัวแปร)ชนิดตัวแปลที่รีเทินกลับไป{ทำงาน}
// ทำการคำนวณค่า แล้วรีเทินหลับไป
func plusReturn(value1 int, value2 int)int{
	return value1+value2
}

// func รับ และ return ค่า 3 ค่า
// value1, value2 ไม่จำเป็นต้องใส่ type int ก็ได้ เพราะ go จะทำการตรวจสอบค่าที่รับเข้ามา ว่าเป็นประเภทไหน ก่อนทำการ บวก
func plus3Return(value1, value2, value3 int) int{
	return value1 + value2 +value3
}

func main() {
	hello()
	plus(1,2)
	result2 := plusReturn(3,4)
	result3 := plus3Return(1,2,3)
	fmt.Println("result = ", result2)
	fmt.Println("result3 =", result3)
}