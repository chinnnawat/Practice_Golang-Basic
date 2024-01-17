package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(promp string)float64 {

	fmt.Printf("%v", promp)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message,_ := fmt.Printf("%v number only", value)
		panic(message)
	}
	return value
}

func getOperator() string{
	fmt.Printf("(+,-,*,/) : ")
	op,_ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func addValue(value1 , value2 float64) float64{
	return value1+value2	
}

func minusValue(value1 , value2 float64)float64{
	return value1-value2
}

func multiValue(value1, value2 float64)float64{
	return value1*value2
}

func devideValue(value1, value2 float64)float64{
	return value1/value2
}

func main() {
	value1 := getInput("value1 = ")
	oparetor := getOperator()
	value2 := getInput("value2 = ")
	var result float64

	switch oparetor{
	case "+":
		result = addValue(value1, value2)
	case "-":
		result = minusValue(value1, value2)
	case "*":
		result = multiValue(value1, value2)
	case "/":
		result = devideValue(value1, value2)
	default :
		fmt.Println("invalid")
	}

	for {
		oparetorNext := getOperator()
		if oparetorNext == "=" {
			fmt.Println("Total", result)
			break
		} else {
			valueNext := getInput("valueNext = ")
			switch oparetorNext{
			case "+":
				result = addValue(result,valueNext)
			case "-":
				result = minusValue(result,valueNext)
			case "*":
				result = multiValue(result,valueNext)
			case "/":
				result = devideValue(result,valueNext)
			default :
				fmt.Println("invalid")
			}
		}
	}
}