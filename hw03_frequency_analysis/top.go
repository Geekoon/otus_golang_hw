package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"sort"
	"unicode"
)

func unpack(a string) []string {
	var (
		curWord   string
		result    []string
		findSpace bool
	)

	for _, v := range a {

		if unicode.IsSpace(v) {
			findSpace = true
			continue
		}

		if findSpace && curWord != "" {
			result = append(result, curWord)
			curWord = ""
			findSpace = false
		}

		curWord += string(v)
	}
	if curWord != "" {
		result = append(result, curWord)
	}

	return result
}

func Top10(text string) []string {
	s := unpack(text)
	m := make(map[string]int)
	//	fmt.Printf("%q\n", s)
	for _, v := range s {
		m[v] += 1
	}
	//	fmt.Println(m)

	var r []struct {
		w string
		n int
	}
	for k, v := range m {
		a := struct {
			w string
			n int
		}{k, v}
		r = append(r, a)
	}
	//	fmt.Println(r)

	sort.Slice(r, func(i, j int) bool { return r[i].n > r[j].n })
	//	fmt.Println(r)

	//	fmt.Println(strings.Count(text, "он"))
	h := 0
	if len(r) < 10 {
		h = len(r)
	} else {
		h = 10
	}

	var res []string
	for i := 0; i < h; i++ {
		res = append(res, r[i].w)
	}

	return res
}
