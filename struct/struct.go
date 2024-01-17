package main

import "fmt"

type employee struct {
	employeeID   string
	employeeName string
	phoneNum     string
}

func arrayStruct(){
	// ****************Struc + Array******************
	// Array data => ต้องมีการบอกขนาดด้วย
	// 1. กำหนดขนาด => ตัวแปร := [ขนาด array]struc{}
	employeeList := [3]employee{}
	// 2. การใส่ค่าให้ array object => ตัวแปร[ตำแหน่งIndex] = struc{กำหนดค่า}
	employeeList[0] = employee{
		employeeID:   "101",
		employeeName: "Chin1",
		phoneNum:     "080000000",
	}
	employeeList[1] = employee{
		employeeID:   "102",
		employeeName: "Chin2",
		phoneNum:     "080000000",
	}
	employeeList[2] = employee{
		employeeID:   "103",
		employeeName: "Chin3",
		phoneNum:     "080000000",
	}
	fmt.Println("employeeList =", employeeList)
}

func sliceStruct(){
	// ****************Struc + Slice******************
	employeeList := []employee{}
	employee1 := employee{
		employeeID:   "101",
		employeeName: "Chin1",
		phoneNum:     "080000000",
	}
	employee2 := employee{
		employeeID:   "102",
		employeeName: "Chin2",
		phoneNum:     "080000000",
	}
	employee3 := employee{
		employeeID:   "103",
		employeeName: "Chin3",
		phoneNum:     "080000000",
	}

	employeeList = append(employeeList, employee1,employee2,employee3)
	fmt.Println(employeeList)
}


func main() {
	// Array + Struct
	// arrayStruct()

	// Slice + Struct
	sliceStruct()

}