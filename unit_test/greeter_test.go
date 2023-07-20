package greeter

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGreetPositive(t *testing.T) {
	testTable := []struct {
		hour int
		name string
		want string
	}{
		{hour: 6, name: "Gabriel", want: "Good morning Gabriel!"},
		{hour: 11, name: "gabriel", want: "Good morning Gabriel!"},
		{hour: 12, name: "Gabriel Jesus", want: "Hello Gabriel Jesus!"},
		{hour: 17, name: "gabriel jesus", want: "Hello Gabriel Jesus!"},
		{hour: 18, name: "    gabriel jesus    ", want: "Good evening Gabriel Jesus!"},
		{hour: 21, name: "пьер", want: "Good evening Пьер!"},
		{hour: 22, name: "пьер-эмерик эмильяно франсуа обамеянг", want: "Good night Пьер-Эмерик Эмильяно Франсуа Обамеянг!"},
		{hour: 23, name: "дж. р. р. толкин", want: "Good night Дж. Р. Р. Толкин!"},
		{hour: 0, name: "欧阳子真", want: "Good night 欧阳子真!"},
		{hour: 5, name: "!@#$%^:()+_==-Ёё`^&124354", want: "Good night !@#$%^:()+_==-Ёё`^&124354!"},
		{name: "Priscilla", want: "Good night Priscilla!"},
	}

	for i, testCase := range testTable {
		t.Run(fmt.Sprintf("test %d", i+1), func(t *testing.T) {
			res, err := Greet(testCase.name, testCase.hour)
			assert.NoError(t, err)
			assert.Equal(t, testCase.want, res)
		})
	}
}

func TestGreetNegative(t *testing.T) {
	testTable := []struct {
		name string
		hour int
		want string
	}{

		{hour: -1, name: "пьер", want: "Введено неверное время. Введите час в промежутке от 0 по 23 включительно."},
		{hour: 24, name: "пьер", want: "Введено неверное время. Введите час в промежутке от 0 по 23 включительно."},
		{hour: 5, name: " ", want: "Введено пустое поле name."},
		{hour: 5, name: "          ", want: "Введено пустое поле name."},
		{hour: 5, want: "Введено пустое поле name."},
	}
	for i, testCase := range testTable {
		t.Run(fmt.Sprintf("test %d", i+1), func(t *testing.T) {
			_, err := Greet(testCase.name, testCase.hour)
			assert.EqualError(t, err, testCase.want)
		})
	}
}
