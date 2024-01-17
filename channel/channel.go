package main

import (
	"fmt"
	"time"
)

// channel ใช้สื่อสารกันระหว่าง routine1 and routine2

// process1(c chan string) => function process1 ที่มี c เป็น channel ซึ่งมี type เป็น string => ท่อ
// (c chan string, data string) => มี c เป็นท่อ ซึ่งส่งข้อมูล data ที่เป็น string มา
func process1(c chan string, data string){
	// ยัด data ที่รับมา ลงใน channel c
	c <- data
}

func main() {
	// ในการสร้างท่อ หรือ channel ว่างๆขึ้นมา จะใช้คำสั่ง make ในการสร้าง

	// ***********Buffer***************
	// make(chan type, จองbuffer)
	// ch := make(chan string,1)

	// เรียกใช้
	// ส่ง (ch,"Data") => (channel string, data string) ไปยัง func process1
	// process1(ch,"Data1")

	// **********Routine****************
	ch := make(chan string)
	go process1(ch,"Datal1")
	fmt.Println(<-ch)
	time.Sleep(5*time.Second)

}