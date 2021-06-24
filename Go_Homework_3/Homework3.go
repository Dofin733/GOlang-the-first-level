package main

import (
	"fmt"
	"math"
	"os"
)

func main(){
	var a,b, r float64
	var op string

	fmt.Print("Введите первое число: ")
	if _, err := fmt.Scanln(&a); err != nil {
		return
	}

	fmt.Print("Введите второе число: ")
	if _, err := fmt.Scanln(&b); err != nil {
		return
	}

	fmt.Print("Введите арифметическую операцию (+, -, *, /, %): ")
	if _, err := fmt.Scanln(&op); err != nil {
		return
	}

	switch op{

	case "+":
		r = a + b

	case "-":
		r = a - b

	case "*":
		r = a * b

	case "/":
		r = a / b

	case "%": //взятие остатка
		r = math.Mod(a,b)

	default: // при введение операции отличной от заданной выдается default
		fmt.Println("Операция выбрана неверно, выберите операцию из списка")
		os.Exit(1)
	}

	//fmt.Printf("Результат выполнения операции: %f\n", r)
	fmt.Print( a, op, b , "=", r)
}