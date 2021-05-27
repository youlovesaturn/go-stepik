package main

import "fmt"

func main() {
	var d int
	var h int
	var m int

	fmt.Scan(&d)
	h = d / 30
	m = 2 * (d % 30)
	fmt.Println("It is", h, "hours", m, "minutes.")
}

// Часовая стрелка повернулась с начала суток на d градусов.
// Определите, сколько сейчас целых часов h и целых минут m.
