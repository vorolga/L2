package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	sort2 "sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки //nkrub
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца //Mkrub
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные //c
-h — сортировать по числовому значению с учётом суффиксов //hkrub

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

const outPutFile = "sorted"

var Months = map[string]int{
	"january":   1,
	"february":  2,
	"march":     3,
	"april":     4,
	"may":       5,
	"june":      6,
	"july":      7,
	"august":    8,
	"september": 9,
	"october":   10,
	"november":  11,
	"december":  12,
	"jan":       1,
	"feb":       2,
	"mar":       3,
	"apr":       4,
	"aug":       8,
	"sept":      9,
	"oct":       10,
	"nov":       11,
	"dec":       12,
}

var Suf = map[string]int{
	"femto": 1,
	"f":     1,
	"pico":  2,
	"p":     2,
	"nano":  3,
	"n":     3,
	"micro": 4,
	"µ":     4,
	"milli": 5,
	"m":     5,
	"centi": 6,
	"c":     6,
	"deci":  7,
	"d":     7,
	"deca":  8,
	"da":    8,
	"hecto": 9,
	"h":     9,
	"kilo":  10,
	"k":     10,
	"mega":  11,
	"M":     11,
	"giga":  12,
	"G":     12,
	"tera":  13,
	"T":     13,
	"peta":  14,
	"P":     14,
}

type Options struct {
	k int
	n bool
	r bool
	u bool
	M bool
	b bool
	c bool
	h bool
}

func b2i(boolValue bool) int {
	if boolValue {
		return 1
	}

	return 0
}

func getOptions() *Options {
	var flags Options
	flag.IntVar(&flags.k, "k", -1, "указание колонки для сортировки")
	flag.BoolVar(&flags.n, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&flags.r, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&flags.u, "u", false, "не выводить повторяющиеся строки")
	flag.BoolVar(&flags.M, "M", false, "сортировать по названию месяца")
	flag.BoolVar(&flags.b, "b", false, "игнорировать хвостовые пробелы")
	flag.BoolVar(&flags.c, "c", false, "проверять отсортированы ли данные")
	flag.BoolVar(&flags.h, "h", false, "сортировать по числовому значению с учётом суффиксов")

	flag.Parse()

	if b2i(flags.n)+b2i(flags.M)+b2i(flags.h) > 1 { //может быть либо обычная сортировка,
		// либо по чиловому значению, либо по названию месяца, либо по числовому значению с учетом суффиксов
		flags.n = false
		flags.M = false
		flags.h = false
	}

	return &flags
}

func readInput(filepath string) ([]string, error) {
	data := make([]string, 0)

	var in io.Reader

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			log.Fatalf("Error closing file.txt: %s", err)
		}
	}(file)

	in = file

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func writeOutput(data []string, filepath string) error {
	var out io.Writer

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			log.Fatalf("Error closing file.txt: %s", err)
		}
	}(file)

	out = file

	for _, str := range data {
		if _, err := io.WriteString(out, str+"\n"); err != nil {
			return err
		}
	}

	return nil
}

func sort(data []string, options Options) ([]string, error) {
	switch {
	case options.c:
		fmt.Println(checkSort(data, options))
		return data, nil
	case options.n:
		switch {
		case options.k >= 0: //сортировка по числовому значению в колонке
			return columnNumberSort(data, options)
		default: //сортировка по числовому значению
			return numberSort(data, options)
		}
	case options.M:
		switch {
		case options.k >= 0: //по месяцу в колоке
			return columnMonthSort(data, options)
		default: //по месяцу
			return monthSort(data, options)
		}
	case options.h:
		switch {
		case options.k >= 0: //по числовому с суффиксами в колонке
			return columnSufSort(data, options)
		default: //по числовому с суффиксами
			return sufSort(data, options)
		}
	default:
		switch {
		case options.k >= 0: //обычная в колонке
			return columnBaseSort(data, options)
		default: //обычная
			return baseSort(data, options)
		}
	}
}

func columnSufSort(data []string, options Options) ([]string, error) {
	if len(data) == 0 {
		return []string{}, nil
	}

	col := options.k

	var er error

	sort2.Slice(data, func(i, j int) bool {
		wordsI := strings.Fields(data[i])
		wordsJ := strings.Fields(data[j])

		if len(wordsI) <= col || len(wordsJ) <= col {
			return false
		}

		if options.b {
			wordsI[col] = strings.TrimRight(wordsI[col], " ")
			wordsJ[col] = strings.TrimRight(wordsJ[col], " ")
		}

		var k int
		for k = 0; k < len(wordsI[col]) && (string(wordsI[col][k]) >= "1" && string(wordsI[col][k]) <= "9"); k++ {
		}

		numI, err := strconv.ParseFloat(wordsI[col][:k], 64)
		if err != nil {
			er = err
		}

		for k = 0; k < len(wordsJ[col]) && (string(wordsJ[col][k]) >= "1" && string(wordsJ[col][k]) <= "9"); k++ {
		}
		numJ, err := strconv.ParseFloat(wordsJ[col][:k], 64)
		if err != nil {
			er = err
		}

		sufI, ok := Suf[wordsI[col][k:]]
		if !ok {
			er = errors.New("there is no suf")
		}

		sufJ, ok := Suf[wordsJ[col][k:]]
		if !ok {
			er = errors.New("there is no suf")
		}

		if options.r {
			return sufI > sufJ || (sufI == sufJ && numI > numJ)
		}

		return sufI < sufJ || (sufI == sufJ && numI < numJ)
	})

	if options.u {
		result := make([]string, 0)
		result = append(result, data[0])

		for i := 1; i < len(data); i++ {
			if data[i] != result[len(result)-1] {
				result = append(result, data[i])
			}
		}

		return result, er
	}

	return data, er
}

func sufSort(data []string, options Options) ([]string, error) {
	var er error

	sort2.Slice(data, func(i, j int) bool {
		var wordsI, wordsJ string
		wordsI = data[i]
		wordsJ = data[j]

		if options.b {
			wordsI = strings.TrimRight(data[i], " ")
			wordsJ = strings.TrimRight(data[j], " ")
		}

		var k int
		for k = 0; k < len(wordsI) && (string(wordsI[k]) >= "1" && string(wordsI[k]) <= "9"); k++ {
		}

		numI, err := strconv.ParseFloat(wordsI[:k], 64)
		if err != nil {
			er = err
		}

		for k = 0; k < len(wordsJ) && (string(wordsJ[k]) >= "1" && string(wordsJ[k]) <= "9"); k++ {
		}
		numJ, err := strconv.ParseFloat(wordsJ[:k], 64)
		if err != nil {
			er = err
		}

		sufI, ok := Suf[wordsI[k:]]
		if !ok {
			er = errors.New("there is no suf")
		}

		sufJ, ok := Suf[wordsJ[k:]]
		if !ok {
			er = errors.New("there is no suf")
		}

		if options.r {
			return sufI > sufJ || (sufI == sufJ && numI > numJ)
		}

		return sufI < sufJ || (sufI == sufJ && numI < numJ)
	})

	return data, er
}

func columnMonthSort(data []string, options Options) ([]string, error) {
	if len(data) == 0 {
		return []string{}, nil
	}

	col := options.k

	var er error

	sort2.Slice(data, func(i, j int) bool {
		wordsI := strings.Fields(data[i])
		wordsJ := strings.Fields(data[j])

		if len(wordsI) <= col || len(wordsJ) <= col {
			return false
		}

		if options.b {
			wordsI[col] = strings.TrimRight(wordsI[col], " ")
			wordsJ[col] = strings.TrimRight(wordsJ[col], " ")
		}

		monthI, ok := Months[strings.ToLower(wordsI[col])]
		if !ok {
			er = errors.New("there is no month")
		}
		monthJ, ok := Months[strings.ToLower(wordsJ[col])]
		if !ok {
			er = errors.New("there is no month")
		}

		if options.r {
			return monthI > monthJ
		}

		return monthI < monthJ
	})

	if options.u {
		result := make([]string, 0)
		result = append(result, data[0])

		for i := 1; i < len(data); i++ {
			if data[i] != result[len(result)-1] {
				result = append(result, data[i])
			}
		}

		return result, er
	}

	return data, er
}

func monthSort(data []string, options Options) ([]string, error) {
	var er error

	sort2.Slice(data, func(i, j int) bool {
		var wordsI, wordsJ string
		wordsI = data[i]
		wordsJ = data[j]

		if options.b {
			wordsI = strings.TrimRight(data[i], " ")
			wordsJ = strings.TrimRight(data[j], " ")
		}

		monthI, ok := Months[strings.ToLower(wordsI)]
		if !ok {
			er = errors.New("there is no month")
		}
		monthJ, ok := Months[strings.ToLower(wordsJ)]
		if !ok {
			er = errors.New("there is no month")
		}

		if options.r {
			return monthI > monthJ
		}

		return monthI < monthJ
	})

	if options.u {
		result := make([]string, 0)
		result = append(result, data[0])

		for i := 1; i < len(data); i++ {
			if data[i] != result[len(result)-1] {
				result = append(result, data[i])
			}
		}

		return result, er
	}

	return data, er
}

func columnNumberSort(data []string, options Options) ([]string, error) {
	col := options.k

	var er error

	sort2.Slice(data, func(i, j int) bool {
		wordsI := strings.Fields(data[i])
		wordsJ := strings.Fields(data[j])

		if len(wordsI) <= col || len(wordsJ) <= col {
			return false
		}

		if options.b {
			wordsI[col] = strings.TrimRight(wordsI[col], " ")
			wordsJ[col] = strings.TrimRight(wordsJ[col], " ")
		}

		numI, err := strconv.ParseFloat(wordsI[col], 64)
		if err != nil {
			er = err
		}
		numJ, err := strconv.ParseFloat(wordsJ[col], 64)
		if err != nil {
			er = err
		}

		if options.r {
			return numI > numJ
		}

		return numI < numJ
	})

	if options.u {
		result := make([]string, 0)
		result = append(result, data[0])

		for i := 1; i < len(data); i++ {
			if data[i] != result[len(result)-1] {
				result = append(result, data[i])
			}
		}

		return result, er
	}

	return data, er
}

func numberSort(data []string, options Options) ([]string, error) {
	var er error

	sort2.Slice(data, func(i, j int) bool {
		var wordsI, wordsJ string
		wordsI = data[i]
		wordsJ = data[j]

		if options.b {
			wordsI = strings.TrimRight(data[i], " ")
			wordsJ = strings.TrimRight(data[j], " ")
		}

		numI, err := strconv.ParseFloat(wordsI, 64)
		if err != nil {
			er = err
		}
		numJ, err := strconv.ParseFloat(wordsJ, 64)
		if err != nil {
			er = err
		}

		if options.r {
			return numI > numJ
		}

		return numI < numJ
	})

	if options.u {
		result := make([]string, 0)
		result = append(result, data[0])

		for i := 1; i < len(data); i++ {
			if data[i] != result[len(result)-1] {
				result = append(result, data[i])
			}
		}

		return result, er
	}

	return data, er
}

func columnBaseSort(data []string, options Options) ([]string, error) {
	col := options.k

	var er error

	sort2.Slice(data, func(i, j int) bool {
		wordsI := strings.Fields(data[i])
		wordsJ := strings.Fields(data[j])

		if len(wordsI) <= col || len(wordsJ) <= col {
			return false
		}

		if options.b {
			wordsI[col] = strings.TrimRight(wordsI[col], " ")
			wordsJ[col] = strings.TrimRight(wordsJ[col], " ")
		}

		if options.r {
			return wordsI[col] > wordsJ[col]
		}

		return wordsI[col] < wordsJ[col]
	})

	if options.u {
		result := make([]string, 0)
		result = append(result, data[0])

		for i := 1; i < len(data); i++ {
			if data[i] != result[len(result)-1] {
				result = append(result, data[i])
			}
		}

		return result, er
	}

	return data, er
}

func baseSort(data []string, options Options) ([]string, error) {
	var er error

	sort2.Slice(data, func(i, j int) bool {
		var wordsI, wordsJ string
		wordsI = data[i]
		wordsJ = data[j]

		if options.b {
			wordsI = strings.TrimRight(data[i], " ")
			wordsJ = strings.TrimRight(data[j], " ")
		}

		if options.r {
			return wordsI > wordsJ
		}

		return wordsI < wordsJ
	})

	if options.u {
		result := make([]string, 0)
		result = append(result, data[0])

		for i := 1; i < len(data); i++ {
			if data[i] != result[len(result)-1] {
				result = append(result, data[i])
			}
		}

		return result, er
	}

	return data, er
}

func checkSort(data []string, options Options) bool {
	return sort2.SliceIsSorted(data, func(i, j int) bool {
		return data[i] < data[j]
	})
}

func main() {
	options := getOptions()

	inputStrings, err := readInput(flag.Arg(0))
	if err != nil {
		log.Fatalf("Input read failed: %s", err)
	}

	sortedStrings, err := sort(inputStrings, *options)
	if err != nil {
		log.Fatalf("Uniq failed: %s", err)
	}

	if err = writeOutput(sortedStrings, outPutFile); err != nil {
		log.Fatalf("Output write failed: %s", err)
	}

}
