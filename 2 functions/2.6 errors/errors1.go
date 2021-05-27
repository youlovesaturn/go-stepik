func main() {
	var a, b int
	fmt.Scan(&a, &b)
	i, err := divide(a, b)
	if err != nil {
		fmt.Print("ошибка")
	} else {
		fmt.Print(i)
	}
}

// Вы должны вызвать функцию divide которая делит два числа, но она возвращает не только результат деления,
// но и информацию об ошибке. В случае какой-либо ошибки вы должны вывести "ошибка", если ошибки нет - результат функции.
// Функция divide(a int, b int) (int, error) получает на вход два числа которые нужно поделить и возвращает результат (int) и ошибку (error).
// Пакет main уже объявлен, а нужные пакеты уже импортированы!
