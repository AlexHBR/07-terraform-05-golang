package main

import (
	"fmt"
	"sort"
)

func main() {
	calc_fut_metr()
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	fmt.Println("2. Наименьший элемент в любом заданном списке \n Список по умолчанию", x)
	fmt.Println("Вариант решения 1 наименьший элемент в любом заданном списке: ", max_mas_no_sort(x))
	fmt.Println("Вариант решения 2 наименьший элемент в любом заданном списке: ", sort_mas(x))
	delitel_3()
}

func calc_fut_metr() {
	fmt.Print("1. Введите  количество метров для перевода в футы: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 0.3048
	fmt.Println("соответсвует", output, " футов \n")
}

func sort_mas(x []int) int {
	sort.Ints(x)
	return x[len(x)-1]
}

func max_mas_no_sort(x []int) int {
	e_max := 0
	for _, i := range x {
		if e_max < i {
			e_max = i
		}
	}
	return e_max
}

func delitel_3() {
	fmt.Print("\n3.Числа от 1 до 100 делящиеся на 3 \n")
	for a := 3; a <= 100; a = a + 3 {
		fmt.Print(" ", a)
	}
	fmt.Println()
}
