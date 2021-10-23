package main

import "fmt"

func main() {
	pro := "Макс - программист"
	a := "автоматизатор"
	w := "c"
	y := 2022
	yy := "года"

	fmt.Println(pro, a, w, y, yy)
	fmt.Printf("%T %T %T %T %T", pro, a, w, y, yy) // функция выводит типы
}
