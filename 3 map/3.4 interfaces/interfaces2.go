package main

import (
	"fmt"
)

type Battery struct {
	Space string
	X     string
}

func (b Battery) String() string {
	return fmt.Sprintf("[%s%s]", b.Space, b.X)
}

func main() {
	var str, space, x string
	fmt.Scan(&str)
	for _, k := range str {
		if k == 48 {
			space += " "
		} else {
			x += "X"
		}
	}

	batteryForTest := Battery{
		Space: space,
		X:     x,
	}
	fmt.Print(batteryForTest)

// Во-первых, вы должны объявить новый тип, удовлетворяющий интерфейсу fmt.Stringer.
// Ваш тип должен предусматривать, что на печати он будет выглядеть так: [      XXXX]: 
// где пробелы - "опустошенная" емкость батареи, а X - "заряженная".
// Во-вторых, на стандартный ввод вы получаете строку, состоящую ровно из 10 цифр: 0 или 1 (порядок 0/1 случайный). 
// Ваша задача считать эту строку любым возможным способом и создать на основе этой строки объект объявленного вами 
// на первом этапе типа: надеюсь, вы понимаете, что строка символизирует емкость батарейки: 
// 0 - это "опустошенная" часть, а 1 - "заряженная".
// В-третьих, созданный вами объект должен называться batteryForTest (использование этого имени обязательно).
