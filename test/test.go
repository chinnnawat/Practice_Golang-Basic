package main

import (
	"fmt"
	"strings"
)

func main() {
    sentence := "   This is a sentence with spaces around it.   "
    
    // ใช้ TrimSpace เพื่อลบช่องว่างที่ขอบเขตของ string
    trimmedSentence := strings.TrimSpace(sentence)

    fmt.Printf("Original sentence: \"%s\"\n", sentence)
    fmt.Printf("Trimmed sentence: \"%s\"\n", trimmedSentence)
}
