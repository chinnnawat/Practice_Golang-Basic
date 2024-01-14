package main

import "fmt"

func setZeroReturn(value int)int{
	value = 0
	return value
}

func setZero(value int){
	value = 0
}

// pointer
// ใช้เพื่อเปลี่ยนค่า ให้ไม่เหมือนกับที่ถูกบันทึกไว้ในระบบ
// การส่งค่า ต้องใช้ '&' => &ตัวแปลที่จะเปลี่ยน
// * => ชี้ว่าเป็น pointer
// func ชื่อฟังก์(ตัวแปลที่รับค่ามา *type(ตัวแปลที่รับเข้ามา)){*ตัวแปลที่รับค่ามา = 0}
func zeropointer(i_pointer *int){
	*i_pointer = 0
}

func main() {
	i := 1
	fmt.Println("i =", i)

	// setZeroReturn
	value := setZeroReturn(i)
	fmt.Println("return setzeroReturn =", value)

	// setZero
	// ค่าไม่เปลี่ยน เพราะ setZero ทำงานใน "pass by value" 
	// ทำให้การเปลี่ยนแปลงค่าของ value ภายใน setZero ไม่มีผลกับ i ที่ถูกส่งเข้าไป.
	setZero(i)
	fmt.Println("i from func setZero=", i)

	// pointer
	// &i => ระบุ address ของค่า i ที่ถูกเก็บไว้
	zeropointer(&i)
	fmt.Println("i from func zeropointer", i)
	fmt.Println("address i from func zeropointer", &i)
}