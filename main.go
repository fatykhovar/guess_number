package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	const min = 1
	const max = 100
	try_count := 1
	var input_number int

	fmt.Println("Игра 'Угадай число' - от 1 до 100 началась!")
	fmt.Println("Угадайте число за 10 попыток!")

	rand.NewSource(time.Now().UnixNano())
	random_number := rand.Intn(max-min+1) + min
	
	for {
		if try_count == 10 {
			fmt.Println("Попытки закончились!")
			break
		}
		fmt.Println("Попытка №", try_count, ". Введите число: ")
		fmt.Scanln(&input_number)
		try_count++
		if random_number == input_number {
			fmt.Println("Поздравляю! Вы угадали число", input_number)
			break
		} else if random_number < input_number{
			fmt.Println("Секретное число меньше👇")
		} else {
			fmt.Println("Секретное число больше👆")
		}
	}

	fmt.Println("Игра закончена!")
}