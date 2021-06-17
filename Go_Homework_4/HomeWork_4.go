package main

import (
	"fmt"
	"sort"
)

func main () {
	//slice := []int{5,9,87,3,27,65,1,8,103,4} //первый вариант слайса
	slice := []int{}
	slice = append(slice,125,5,6,3,1,25,72)//второй вариант
	fmt.Println("len:", len(slice), "cap:", cap(slice), "slice:", slice)//вывод данных о слайсе
	sort.Ints(slice)//сортировка слайса
	fmt.Println(slice)//вывод результата сортировки
}
