package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a > 0 {
		fmt.Print("Число положительное")
	} else if a < 0 {
		fmt.Print("Число отрицательное")
	} else if a == 0 {
		fmt.Print("Ноль")
	}
}

// На ввод подается целое число. Если число положительное - вывести сообщение "Число положительное",
// если число отрицательное - "Число отрицательное". Если подается ноль - вывести сообщение "Ноль".
// Выводить сообщение без кавычек.
