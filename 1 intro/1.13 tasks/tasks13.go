package main

import "fmt"

func main() {
	var n int
	f1 := 1
	f2 := 1
	res := -1
	i := 3
	fmt.Scan(&n)
	for ; f2 < n; i++ {
		f1, f2 = f2, f1+f2
		if f2 == n {
			res = i
		}
	}
	fmt.Print(res)
}

// Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является,
// то есть выведите такое число n, что φn=A. Если А не является числом Фибоначчи, выведите число -1.
