package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Options struct {
	f int
	d string
	s bool
}

func getOptions() *Options {
	var flags Options
	flag.IntVar(&flags.f, "f", 0, "\"fields\" - выбрать поля (колонки)")
	flag.StringVar(&flags.d, "d", "\t", "\"delimiter\" - использовать другой разделитель")
	flag.BoolVar(&flags.s, "s", false, "\"separated\" - только строки с разделителем")

	flag.Parse()

	return &flags
}

func readInput(filepath string) ([]string, error) {
	data := make([]string, 0)

	var in io.Reader
	switch len(filepath) {
	case 0:
		in = os.Stdin
	default:
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
	}

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func cut(options *Options, inputStrings []string) []string {
	var result []string
	for _, val := range inputStrings {
		if options.s && !strings.Contains(val, options.d) {
			continue
		}

		columns := strings.Split(val, options.d)
		if len(columns) == 1 {
			result = append(result, val)
			continue
		}

		if len(columns) < options.f {
			result = append(result, "")
			continue
		}

		result = append(result, columns[options.f-1])
	}

	return result
}

func main() {
	options := getOptions()

	if options.f <= 0 {
		log.Fatal("usage: -f list [-s] [-d] [file.txt ...]\nf > 0")
	}

	inputStrings, err := readInput(flag.Arg(0))
	if err != nil {
		log.Fatalf("Input read failed: %s", err)
	}

	result := cut(options, inputStrings)

	for _, val := range result {
		fmt.Println(val)
	}
}
