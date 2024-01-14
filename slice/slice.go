package main

import "fmt"

func main() {
	var courseName = []string{"javascript", "python"}

	// ********* เพิ่มข้อมูล ***********
	// append => ใช้เพิ่มข้อมูล
	// appemd( ตัวแปรที่เราจเพิ่ม , ค่าที่เราจะใส่เพิ่มเข้าไป )
	// append สามารถเพิ่มข้อมูลหลายๆตัวทีเดียวได้ เช่น courseNameUpdate
	courseName = append(courseName,"Golang")
	var courseNameUpdate[]string = append(courseName, "HTML", "CSS", "C", "C++")
	fmt.Println(courseName)
	fmt.Println(courseNameUpdate)

	// ********* ลบข้อมูล ***********
	// [4:7] จะลบข้อมูลตต 1-4 และจะเก็บข้อมูล ตต 5-7
	courseNameDelete47 := courseNameUpdate[4:7]
	fmt.Println(courseNameDelete47)
	// [:4] เอาข้อมูลที่ 1-4
	courseNameDelete14 := courseNameUpdate[:4]
	fmt.Println(courseNameDelete14)
}