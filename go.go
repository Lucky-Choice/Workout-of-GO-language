package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for i := 0; i <= 7; i++ { // упрошенный вариант первого цикла
		fmt.Println(i)
	}

	text := []string{"this", "is", "learn", "Go", "for", "me"}
	fmt.Println(text[3])     // вывесть конкретное значение из массива по его индексу
	for _, i := range text { // вывести все значения массива по порядку
		fmt.Println(i)
	}

	text1 := []string{"this", "is", "learn", "Go", "for", "me"}
	for _, i := range text1 {
		if i == "Go" { // если значение равно "Go" то цикл заканчивается. Значение Go не выведится т.к. i == "Go"
			break
		}
		fmt.Println(i)
	}

}
