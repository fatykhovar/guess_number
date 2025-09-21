package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Statistics struct {
	Date     time.Time `json:"date"`
	IsWin    bool      `json:"isWin"`
	TryCount int       `json:"tryCount"`
}

func main() {
	var results []Statistics
	file, _ := os.Create("results.json")
	defer file.Close()

	for {
		fmt.Println("Игра 'Угадай число' - от 1 до 100 началась!")
		fmt.Println("Угадайте число за 10 попыток!")

		min_number, max_number, try_limit := chooseLevel()
		random_number := generateNumber(min_number, max_number)

		try_counter := 1
		input_numbers := make([]int, 0, try_limit)
		isWin := false

		for {
			input_number, status := inputNumber(try_limit, &try_counter, random_number)

			if status == "continue" {
				continue
			} else if status == "break" {
				break
			}

			input_numbers = append(input_numbers, input_number)

			isBreak := checkInputNumber(random_number, input_number)
			if isBreak {
				isWin = true
				break
			}

			fmt.Printf("Предыдущие попытки: %v\n", input_numbers)
		}
		fmt.Println("Игра закончена!")

		results = append(results, Statistics{Date: time.Now(), IsWin: isWin, TryCount: try_counter})
		
		if isRetry := askRetry(); !isRetry {
			json.NewEncoder(file).Encode(results)
			break
		}
		fmt.Println("isRetry")
	}
}

func chooseLevel() (min int, max int, try int) {
	color.Magenta("Выбери уровень игры:")
	fmt.Println("1. Easy: 1–50, 15 попыток",
		"\n2. Medium: 1–100, 10 попыток",
		"\n3. Hard: 1–200, 5 попыток.",
		"\nВведи номер уровня:")
	var level int
	for {
		_, err := fmt.Scanln(&level)
		if err != nil{
			var trash string
			fmt.Scanln(&trash)
			color.Red("Ошибка ввода! Нужно ввести число 1-3.")
			continue
		}

		if  level < 1 || level > 3{
			color.Red("Ошибка ввода! Нужно ввести число 1-3.")
			continue
		}

		break

	}

	switch level {
	case 1:
		return 1, 50, 15
	case 2:
		return 1, 100, 10
	case 3:
		return 1, 200, 5
	default:
		return 0, 0, 0
	}

}

func generateNumber(min int, max int) int {
	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(max - min + 1) + min
}

func inputNumber(try_limit int, try_counter *int, random_number int) (input int, status string) {

	var input_number int
	if *try_counter > try_limit {
		color.Red("Попытки закончились! Загаданное число: %d", random_number)
		return 0, "break"
	}

	color.Yellow("Попытка №%d. Введите число: ", *try_counter)

	_, err := fmt.Scanln(&input_number)
	if err != nil {
		var trash string
		fmt.Scanln(&trash)

		color.Red("Ошибка ввода! Нужно ввести число.\n")
		return 0, "continue"
	}

	*try_counter++
	return input_number, ""
}

func checkInputNumber(random_number int, input_number int) (isBreak bool) {
	var prompt string
	abs := math.Abs(float64(random_number - input_number))
	if abs <= 5 {
		prompt += "🔥 Горячо!"
	} else if abs <= 15 {
		prompt += "🙂 Тепло!"
	} else {
		prompt += "❄️ Холодно!"
	}

	if random_number == input_number {
		color.Green("Поздравляю! Вы угадали число %d", input_number)
		return true
	} else if random_number < input_number {
		prompt += " Секретное число меньше👇"
	} else {
		prompt += " Секретное число больше👆"
	}

	fmt.Println(prompt)
	return false
}

func askRetry() bool {
	color.Cyan("\nСыграть еще раз? (Да/Нет)")
	var retry string
	for {
		_, err := fmt.Scanln(&retry)
		if err != nil {
			var trash string
			fmt.Scanln(&trash)
			color.Red("Ошибка ввода! Нужно ввести ответ Да/Нет.")
			continue
		}
		if strings.ToLower(retry) != "да" && strings.ToLower(retry) != "нет"{
			color.Red("Ошибка ввода! Нужно ввести ответ Да/Нет.")
			continue
		}
		break
	}
	return strings.ToLower(retry) == "да"
}
