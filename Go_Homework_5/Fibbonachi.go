package main

import "fmt"

func fibbonachi(n int) int{

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibbonachi(n - 1) + fibbonachi(n - 2)
}
func main() {
	//m := map[int]int{}// так и не смог передать значения в мапу, не понимаю как реализовать
	//m[1] = 1
	//m[2] = 1
	fmt.Println(fibbonachi(4)) // 3
	fmt.Println(fibbonachi(5)) // 5
	fmt.Println(fibbonachi(6)) // 8
	fmt.Println(fibbonachi(7)) //13
	fmt.Println(fibbonachi(8)) //21
}


