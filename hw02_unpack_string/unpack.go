package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
)

var ErrInvalidString = errors.New("invalid string")

// Функция распаковки строки по указанному количеству повторений символов.
func Unpack(a string) (string, error) {
	var result string = ""
	var takeLetter bool = true

	for id, v := range a {
		if id == len(a)-1 {
			// Обрабатываем последний символ в строке
			if (v > 47) && (v < 58) {
				return "", ErrInvalidString
			} else {
				result += string(v)
			}
		} else {
			if takeLetter {
				if (v > 47) && (v < 58) {
					// Ошибка, если число вместо символа
					return "", ErrInvalidString
				} else {
					// Если есть повторения, то узнаем их количество и записываем в результат
					if (a[id+1] > 46) && (a[id+1] < 58) {
						j, _ := strconv.Atoi(string(a[id+1]))
						for k := 0; k < j; k++ {
							result += string(v)
						}
						takeLetter = false
					} else {
						// Записываем одиночный символ
						result += string(v)
					}
				}
			} else {
				// Пропускаем цифру
				takeLetter = true
			}
		}
	}
	return result, nil
}
