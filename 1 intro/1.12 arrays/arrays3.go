package main

import "fmt"

func main() {
	arr := [5]int{}
	a := 0
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		arr[i] = a
	}
	for i := 0; i < 5; i++ {
		switch {
		case a < arr[i]:
			a = arr[i]
		}
	}
	fmt.Println(a)

}

// На ввод подаются пять чисел, которые записываются в массив. Однако эта часть программы уже написана.
// Вам нужно написать фрагмент кода, с помощью которого можно найти и вывести максимальное число в этом массиве.
