package main

import (
	"fmt"
	"time"
)

func from(promp string) {
	for i := 0; i < 500; i++ {
		fmt.Println(promp, ":", i)
	}
}

func main() {
	// go => ให้ใช้ routine ของ go ในการทำงาน โดยเป็นการสวิช การทำงานโดยไม่ต้องรอให้อันแรกเสร็จก่อนค่อยทำอันที่2 แต่อันไหนเสร็จก่อนให้ทำอันนั้นก่อนได้เลย
	// ข้อดีคือ สมมุดอันแรก วนลูป 1000 รอบ แต่อันที่ 2 วนแค่ 10 รอบ ก็สามารถให้อันที่ 2 ทำงานก่อนได้เลย โดยไม่ต้องรออันแรกเสร็จ
	go from("Hello")
	go from("Chinnawat")
	time.Sleep(5*time.Second)
}