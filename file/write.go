package main

import "os"

func main() {

	// ข้อมูล data1 จะเป็นรูปแบบ byte
	data1 := []byte("hello \n Chinnawat Wongket")
	err := os.WriteFile("data.txt",data1,0644)

	if err != nil {
		panic(err)
	}

	// Create File
	f,ferr := os.Create("createEmployee.txt")
	if ferr != nil {
		panic(ferr)
	}

	// เขียน ทับลงในไฟล์
	data2 := []byte("hello \n New Employee")
	os.WriteFile("createEmployee.txt", data2, 0644)

	defer f.Close()
}