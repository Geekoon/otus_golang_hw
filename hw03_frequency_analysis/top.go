package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"sort"
	"unicode"
)

// Функция разделения текста на слова по пробелам.
func unpack(a string) []string {
	var (
		result    []string
		curWord   string
		findSpace bool
	)

	for _, v := range a {
		// Прупскаем любые пробелы, табуляции, перводы строки
		if unicode.IsSpace(v) {
			findSpace = true
			continue
		}

		// По найденному пропуску после окончания слова записываем это слово в result
		// Доп проверка на пробел в начале текста
		if findSpace && curWord != "" {
			result = append(result, curWord)
			curWord = ""
			findSpace = false
		}

		// Составляем слово по символьно
		curWord += string(v)
	}

	// Записываем последнее слово в result
	if curWord != "" {
		result = append(result, curWord)
	}
	return result
}

// Функция поиска наиболее повторяющихся слов.
func Top10(text string) []string {
	sepText := unpack(text)

	// Записываем все слова без повторений и считаем количество
	countWords := make(map[string]int)
	for _, v := range sepText {
		countWords[v]++
	}

	r := []struct {
		w string
		n int
	}{}

	// Переносим слова и их количество в структуру
	for k, v := range countWords {
		r = append(r, struct {
			w string
			n int
		}{k, v})
	}

	// Сортируем слова по их количеству
	sort.Slice(r, func(i, j int) bool { return r[i].n > r[j].n })

	// Задаем максимальное число слов для возврата
	h := 0
	if len(r) < 10 {
		h = len(r)
	} else {
		h = 10
	}

	// Заносим нужное количество слов для возрата из функции
	var res []string
	for i := 0; i < h; i++ {
		res = append(res, r[i].w)
	}

	return res
}
