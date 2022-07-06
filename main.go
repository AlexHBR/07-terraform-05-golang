package main

import (
	"fmt"
	"sort"
)

func main() {
	//sort_mas()
	max_mas_no_sort()
	//calc_fut_metr()

	//delitel_3()

}

func calc_fut_metr() {

	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 0.3048

	fmt.Println(output)

}

func sort_mas() {
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	sort.Ints(x)
	fmt.Println(x[len(x)-1])
}

func max_mas_no_sort() {
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	e_max := 0
	for _, s := range x {
		if e_max < s {
			e_max = s
		}
	}

	fmt.Println(e_max)
}

func delitel_3() {

	for a := 3; a <= 100; a = a + 3 {
		fmt.Printf("value of a: %d\n", a)

	}
}
