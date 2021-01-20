package main

import (
	"fmt"
	"strconv"
	"bufio"
	"os"
)

func main() {
	scanNum1 := bufio.NewScanner(os.Stdin)
	fmt.Println("Number 1: ")
	scanNum1.Scan()
	num1 := scanNum1.Text()
	
	scanOperator := bufio.NewScanner(os.Stdin)
	fmt.Println("['+', '-', '*', '/']: ")
	scanOperator.Scan()
	operator := scanOperator.Text()
	
	scanNum2 := bufio.NewScanner(os.Stdin)
	fmt.Println("Number 2: ")
	scanNum2.Scan()
	num2 := scanNum2.Text()

	floatNum1, _ := strconv.ParseFloat(num1, 1)
	floatNum2, _ := strconv.ParseFloat(num2, 1)
	fmt.Println("Num 1",floatNum1, "Num 2", floatNum2)
	
	switch operator {
	case "+":
		fmt.Printf("Answer: %v", (floatNum1+floatNum2))
		break
	case "-":
		fmt.Printf("Answer: %v", (floatNum1-floatNum2))
		break
	case "*":
		fmt.Printf("Answer: %v", (floatNum1*floatNum2))
		break
	case "/":
		fmt.Printf("Answer: %v", (floatNum1/floatNum2))
		break
	default:
		fmt.Println("Unknown Operator.")
	}

}

