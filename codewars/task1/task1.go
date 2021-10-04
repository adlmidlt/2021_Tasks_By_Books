package task1

import (
	"fmt"
	"strings"
)

/* 	Завершите решение так, чтобы оно разбило строку на пары из двух символов.
Если строка содержит нечетное количество символов, она должна заменить отсутствующий
второй символ последней пары подчеркиванием ('_').
*/

func Solution(str string) []string {

	var words []string

	if (len(str) % 2) == 1 {
		str = fmt.Sprintf("%v%v", str, "_")
	}

	a := strings.Split(str, "")

	for i := 0; i < len(a); i++ {
		t := fmt.Sprintf("%v%v", a[i], a[i+1])
		words = append(words, t)
		i++
	}
	fmt.Println(words)
	return words
}
