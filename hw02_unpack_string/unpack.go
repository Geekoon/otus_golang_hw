package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

// Функция распаковки строки по указанному количеству повторений символов.
func Unpack(a string) (string, error) {
	var result string = ""
	var curLetter rune = 0
	var resultErr error = nil

	for _, v := range a {
		//	fmt.Printf("%s = %d\n", string(v), v)

		if unicode.IsDigit(v) && curLetter == 0 { //nolint
			// Ошибка, если число вместо символа
			result = ""
			resultErr = ErrInvalidString
			break
		} else {
			if unicode.IsDigit(v) {
				// Повторяем символ по числу после него
				j, _ := strconv.Atoi(string(v))
				result += strings.Repeat(string(curLetter), j)
				curLetter = 0
			} else {
				// Записываем одиночный символ
				if curLetter != 0 {
					result += string(curLetter)
				}
				curLetter = v // запоминаем символ на случай цифры после него
			}
		}
	}

	// Записываем последний символ, исключая нулевое значение руны
	if !unicode.IsDigit(curLetter) && curLetter != 0 {
		result += string(curLetter)
	}

	return result, resultErr
}
