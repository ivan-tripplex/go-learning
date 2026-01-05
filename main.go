package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, i float64
	var op string

	fmt.Println("Enter first number:")
	fmt.Scan(&x)
	fmt.Println("Enter operation (+,-,/,*,sqrt,**,ln)")
	fmt.Scan(&op)
	fmt.Println("Enter second number:")
	fmt.Scan(&y)
	res := operations(x, y, op)
	fmt.Printf("result: %f\n", res)
	for op != "0" {
		fmt.Println("Enter operation (+,-,/,*,sqrt,**,ln)")
		fmt.Scan(&op)
		if op != "end" {
			fmt.Println("Enter next number:")
			fmt.Scan(&i)
			res = operations(res, i, op)
			fmt.Printf("result: %f\n", res)
		} else {
			break
		}
	}

}

func operations(x float64, y float64, op string) float64 {
	switch op {
	case "+":
		return x + y
	case "-":
		return x - y
	case "*":
		return x * y
	case "/":
		return x / y
	case "sqrt":
		return math.Sqrt(x)
	case "**":
		return math.Pow(x, y)
	case "ln":
		return math.Log(x)
	default:
		fmt.Println("Invalid operation")
		return 0
	}
}
