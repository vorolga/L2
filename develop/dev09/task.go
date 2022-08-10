package main

import (
	"errors"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func downloadFile(filepath string, url string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err = out.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(out)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return errors.New("response status: " + resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Parse()

	if len(flag.Args()) != 1 {
		log.Fatal("Wrong count of args")
	}

	url := flag.Arg(0)
	if !strings.Contains(flag.Arg(0), "://") {
		url = "http://" + url
	}

	strs := strings.Split(url, "/")

	filename := ""
	if len(strs) == 3 {
		filename = "index.html"
	} else {
		filename = strs[len(strs)-1]
	}

	err := downloadFile(filename, url)
	if err != nil {
		err = os.Remove(filename)
		if err != nil {
			log.Fatal(err)
		}
		log.Fatal(err)
	}
}
