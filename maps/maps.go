package main

import "fmt"

func main() {

	// map ที่ไม่มีการกำหนดค่าเริ่มต้น
	// map[key(type)]value(type) => จากตัวอย่าง key => string
	// การใช้งาน map => [key] ตามด้วย value
	// การ map ควรใช้แบบมี make มากกว่า
	var product = make(map[string]float64)
	// product := map[string]string{}

	// add
	// Macbook = key
	product["Macebook"] = 1000
	product["Lenovo"] = 1500
	// add
	// product := map[string]string{"Macbook": "1000", "Acer":"1500"}
	fmt.Println("product =",product)

	// delete
	// product = ตัวแปรที่เก็บ, "Lenovo" = key เมื่อกำหนด key ที่จะลบ จะลบทั้ง key และ value
	delete(product,"Lenovo")
	fmt.Println("productDelete =",product)

	// Update
	product["Iphone"] = 25000
	fmt.Println("productUpdate", product)

	// value
	// การเข้าถึง value ของข้อมูลโดยการใช้ key
	// ตัวแปรที่เก็บ[key]
	value1 := product["Macebook"]
	fmt.Println("Macebook ราคา =", value1)

	// map ที่มีการกำหนดค่าเริ่มต้น
	// map[key(type)value(type){มีข้อมูลอะไรอยู่ในนี้บ้าง}]
	courseName := map[string]string{"101":"JavaScript", "102":"Pythone"}
	fmt.Println("courseName =", courseName)
}