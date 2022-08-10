package main

import (
	"fmt"
	"reflect"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	words := []string{"тяпка", "пятак", "пятка", "листок", "слиток", "столик", "рим", "амамам", "мир", "мама"}
	fmt.Println(anagrams(words))
}

func anagrams(words []string) map[string][]string {
	if len(words) < 1 {
		return map[string][]string{}
	}

	groups := make(map[string][]string)
	groups[words[0]] = append(groups[words[0]], words[0])

	keysChars := make(map[string]map[rune]int)
	for _, ch := range words[0] {
		if keysChars[words[0]] == nil {
			keysChars[words[0]] = map[rune]int{}
		}

		if _, ok := keysChars[words[0]][ch]; !ok {
			keysChars[words[0]][ch] = 1
			continue
		}
		keysChars[words[0]][ch]++
	}

	for word := 1; word < len(words); word++ {
		wordChars := make(map[rune]int)
		for _, ch := range words[word] {
			wordChars[ch]++
		}

		var isAdded bool
	Loop:
		for k := range keysChars {
			if len(reflect.ValueOf(keysChars[k]).MapKeys()) != len(reflect.ValueOf(wordChars).MapKeys()) {
				continue Loop
			}

			for ch := range keysChars[k] {
				if keysChars[k][ch] != wordChars[ch] {
					continue Loop
				}
			}

			groups[k] = append(groups[k], words[word])
			isAdded = true
			break
		}

		if !isAdded {
			groups[words[word]] = []string{words[word]}

			for _, ch := range words[word] {
				if keysChars[words[word]] == nil {
					keysChars[words[word]] = map[rune]int{}
				}
				keysChars[words[word]][ch]++
			}
		}
	}
	return groups
}
