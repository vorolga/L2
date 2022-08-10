package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Options struct {
	A int
	B int
	C int
	c bool
	i bool
	v bool
	F bool
	n bool
}

func getOptions() *Options {
	var flags Options
	flag.IntVar(&flags.A, "A", 0, "'after' печатать +N строк после совпадения")
	flag.IntVar(&flags.B, "B", 0, "'before' печатать +N строк до совпадения")
	flag.IntVar(&flags.C, "C", 0, "'context' печатать ±N строк вокруг совпадения")
	flag.BoolVar(&flags.c, "c", false, "'count' (количество строк)")
	flag.BoolVar(&flags.i, "i", false, "'ignore-case' (игнорировать регистр)")
	flag.BoolVar(&flags.v, "v", false, "'invert' (вместо совпадения, исключать)")
	flag.BoolVar(&flags.F, "F", false, "'fixed', точное совпадение со строкой, не паттерн")
	flag.BoolVar(&flags.n, "n", false, "'line num', печатать номер строки")

	flag.Parse()

	if flags.C > 0 {
		flags.A = flags.C
		flags.B = flags.C
	}

	return &flags
}

func readInputAndFind(filepath string, pattern string, options *Options) ([]string, []int, error) {
	data := make([]string, 0)

	indexes := make([]int, 0)

	var in io.Reader

	file, err := os.Open(filepath)
	if err != nil {
		return nil, nil, err
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			log.Fatalf("Error closing file.txt: %s", err)
		}
	}(file)

	in = file

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		str := scanner.Text()

		cmpStr := str
		if options.i {
			cmpStr = strings.ToLower(cmpStr)
		}

		switch options.F {
		case true:
			if cmpStr == pattern {
				indexes = append(indexes, len(data))
			}
		default:
			if strings.Contains(cmpStr, pattern) {
				indexes = append(indexes, len(data))
			}
		}
		data = append(data, str)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return data, indexes, nil
}

func printStr(options *Options, j int, str string) {
	if options.n {
		fmt.Println(strconv.Itoa(j+1) + ": " + str)
		return
	}

	fmt.Println(str)
}

func grep(options *Options, inputStrings []string, indexes []int) {
	if options.c {
		fmt.Println(len(indexes))
		return
	}

	var lastPrintedIdx int

	for i, val := range indexes {
		if lastPrintedIdx == len(inputStrings) {
			return
		}
		switch options.v {
		case true:
			if i == 0 {
				for j := 0; j < int(math.Min(float64(val+options.A), float64(len(inputStrings)))); j++ {
					printStr(options, j, inputStrings[j])
				}
				lastPrintedIdx = val + options.A
				continue
			}

			if val-1 == indexes[i-1] {
				continue
			}

			for j := int(math.Max(float64(indexes[i-1]+1-options.B), float64(lastPrintedIdx))); j < int(math.Min(float64(val+options.A), float64(len(inputStrings)))); j++ {
				if j != lastPrintedIdx {
					fmt.Println("--")
				}
				printStr(options, j, inputStrings[j])
				lastPrintedIdx = j + 1
			}

		default:
			for j := int(math.Max(float64(lastPrintedIdx), float64(val-options.B))); j < int(math.Min(float64(val+1+options.A), float64(len(inputStrings)))); j++ {
				if j != lastPrintedIdx {
					fmt.Println("--")
				}
				printStr(options, j, inputStrings[j])
				lastPrintedIdx = j + 1
			}

			if i == len(indexes)-1 {
				return
			}
		}
	}
	for j := int(math.Max(float64(indexes[len(indexes)-1]+1-options.B), float64(lastPrintedIdx))); j < len(inputStrings); j++ {
		if j != lastPrintedIdx {
			fmt.Println("--")
		}
		printStr(options, j, inputStrings[j])
		lastPrintedIdx = j + 1
	}
}

func main() {
	options := getOptions()
	inputStrings, indexes, err := readInputAndFind(flag.Arg(1), flag.Arg(0), options)
	if err != nil {
		log.Fatalf("Input read failed: %s", err)
	}
	grep(options, inputStrings, indexes)
}
