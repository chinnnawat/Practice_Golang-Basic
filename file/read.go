package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("indexInfo.csv")
	if err!= nil {
		panic(err)
	}

	// NewScanner ใช้ในการอ่นไฟล์
	scanner := bufio.NewScanner(file)
	
	count := 1
	for scanner.Scan(){
		line := scanner.Text()

		// Split ใช้ในการนำเอาอะไรบางอย่างออกเช่น ","
		// strings.Split(ตัวแปรที่ต้องการนำเอาอะไรออก(string), "สิ่งที่ต้องการเอาออก(string)")
		item := strings.Split(line, ",")

		// print indexตำแหน่งที่ 3
		fmt.Println(item[3])
		count ++
	}
}