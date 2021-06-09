package main

import (
"fmt"
)

func main (){

	var n int

	fmt.Print("Введите трехзначное число: ")
	fmt.Scanln(&n)

	fmt.Println("Количество сотен в числе равно: ",  n/100)
	fmt.Println("Количество десятков в числе равно: ", n/10%10)
	fmt.Println("Количество единиц в числе равно: ", n%10 )
}