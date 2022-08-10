package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func unpack(str string) (string, error) {
	if len(str) == 0 {
		return str, nil
	}

	if unicode.IsDigit(rune(str[0])) {
		return "", errors.New("некорректная строка")
	}

	res := make([]rune, 0)

	for i, v := range []rune(str) {
		j, err := strconv.Atoi(string(v))

		switch {
		case v == '\\':
		case err == nil:
			switch str[i-1] != '\\' {
			case true:
				for k := 1; k < j; k++ {
					res = append(res, rune(str[i-1]))
				}
			case false:
				res = append(res, v)
			}
		default:
			res = append(res, v)
		}
	}

	return string(res), nil
}

func main() {
	var str string

	if _, err := fmt.Scanln(&str); err != nil && !strings.Contains(err.Error(), "unexpected newline") {
		log.Fatal(err)
	}

	un, err := unpack(str)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(un)
}
