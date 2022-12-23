package main

import (
	"fmt"

	"github.com/Venkatesh-8203/CalculatorOp/Add"

	AddOperation "github.com/Venkatesh-8203/CalculatorOp/Add"
	DivOperation "github.com/Venkatesh-8203/CalculatorOp/Division"
	MulOperation "github.com/Venkatesh-8203/CalculatorOp/Multiply"
	SubOperation "github.com/Venkatesh-8203/CalculatorOp/Subtract"
)

func main() {
	a := 100
	b := 50

	addResult := AddOperation.Add(a, b)
	subResult := SubOperation.Subtract(a, b)
	mulResult := MulOperation.Multiply(a, b)
	divResult := DivOperation.Divide(a, b)

	fmt.Println("Add: ", addResult)
	fmt.Println("Sub: ", subResult)
	fmt.Println("Mul: ", mulResult)
	fmt.Println("Div: ", divResult)

	addResult1 := Add.Add(a, b)
	fmt.Println("addResult1: ", addResult1)
}
