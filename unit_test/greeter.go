package greeter

import (
	"errors"
	"fmt"
	"strings"
)

func Greet(name string, hour int) (string, error) {
	greeting := "Good night"
	name = strings.Trim(name, " ")
	if name == "" || name == " " {
		return greeting, errors.New("Введено пустое поле name.")
	}
	if hour < 0 || hour > 23 {
		return greeting, errors.New("Введено неверное время. Введите час в промежутке от 0 по 23 включительно.")
	}

	if hour >= 6 && hour < 12 {
		greeting = "Good morning"
	} else if hour >= 12 && hour < 18 {
		greeting = "Hello"
	} else if hour >= 18 && hour < 22 {
		greeting = "Good evening"
	}
	trimmedName := strings.Trim(name, " ")
	return fmt.Sprintf("%s %s!", greeting, strings.Title(trimmedName)), nil
}
