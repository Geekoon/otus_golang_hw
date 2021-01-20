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
	var (
		curLetter rune // храним обрабатываемый символ для следующей итерации
		result    strings.Builder
	)

	for _, v := range a {
		if unicode.IsDigit(v) && curLetter == 0 {
			// Возвращаем ошибку, если число вместо символа
			return "", ErrInvalidString
		}
		if unicode.IsDigit(v) {
			// Повторяем символ по числу после него
			j, _ := strconv.Atoi(string(v))
			result.WriteString(strings.Repeat(string(curLetter), j))
			curLetter = 0
			continue
		}
		// Записываем одиночный символ
		if curLetter != 0 {
			result.WriteRune(curLetter)
		}
		curLetter = v // запоминаем символ на случай цифры после него
	}

	// Записываем последний символ, исключая нулевое значение руны
	if !unicode.IsDigit(curLetter) && curLetter != 0 {
		result.WriteRune(curLetter)
	}

	return result.String(), nil
}
