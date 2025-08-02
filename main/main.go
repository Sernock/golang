package main

import (
	"fmt"
	"golang/tempconv"
)

func main() {
	c := tempconv.FreezingC // значение типа Celsius
	f := tempconv.CToF(c)   // конвертируем в Fahrenheit
	fmt.Println(f)          // вызывается метод String(), т.к. f имеет тип Fahrenheit
}
