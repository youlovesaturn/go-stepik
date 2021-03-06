package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	x1 := a / 100000
	x2 := a / 10000 % 10
	x3 := a / 1000 % 10
	y1 := a / 100 % 10
	y2 := a / 10 % 10
	y3 := a % 10

	switch {
	case (x1 + x2 + x3) == (y1 + y2 + y3):
		fmt.Print("YES")
	default:
		fmt.Print("NO")
	}
}

// Определите является ли билет счастливым.
// Счастливым считается билет, в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.
