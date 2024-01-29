package main

import "fmt"

func add(value1, value2 float64) {
	result := value1 + value2
	fmt.Println("result : ", result)
}

func loop(){
	for i := 0; i < 10; i++ {
		fmt.Println("i = ", i)
	}
}

// จะแสดงลำดับแบบ (LIFO)
func defer_loop(){
	for i := 0; i < 10; i++ {
		defer fmt.Println("i = ", i)
	}
}

func main() {
	fmt.Println("Welcome to calculate :")

	// defer คือการกำหนดให้ทำงานสิ่งนั้นเป็นลำดับสุดท้าย
	// สามารถใช้ defer หลายตัวใน func เดียวกันได้ และจะเป็นการทำงานแบบ Last in, First Out (LIFO)
	defer fmt.Println("End")
	add(10,20)

	loop()
	defer_loop()
}