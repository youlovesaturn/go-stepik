func T() float64 {
	w := W()
	return 6 / w
}

func W() float64 {
	m := M()
	math.Sqrt((k / m))
}

func M() float64 {
	m := p * v
	return m
}

// Напишите три функции, каждая из которых будет выполнять конкретную формулу.
// Название функций обязательно должны соответствовать букве формулы: T(), W() и M().
// Для того чтобы найти t - необходимо сначала найти w, и т.д.
// Так что используйте результат функции W() в формуле функции T() - то-есть вызывайте функцию W() в T().
// Аналогично и с W(), M(). 
