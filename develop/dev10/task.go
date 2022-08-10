package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

type Options struct {
	timeout int
	host    string
	port    string
}

func getOptions() *Options {
	var flags Options
	flag.IntVar(&flags.timeout, "timeout", 10, "timeout for connect")

	flag.Parse()

	return &flags
}

func connect(options *Options) (*net.Conn, error) {
	connectStr := fmt.Sprintf("%s:%s", options.host, options.port)
	fmt.Println("Trying", connectStr, "...")
	conn, err := net.DialTimeout("tcp", connectStr, time.Duration(options.timeout)*time.Second)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to", connectStr, ".")
	return &conn, nil
}

func sendRequest(conn net.Conn) error {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input := scanner.Bytes()

			_, err := conn.Write(input)
			if err != nil {
				return err
			}
			_, err = conn.Write([]byte("\n"))
			if err != nil {
				return err
			}

			out := make([]byte, 100)

			_, err = conn.Read(out)
			if err != nil {
				return err
			}

			fmt.Println(string(out))
		}
	}
}

func main() {
	options := getOptions()

	if len(flag.Args()) != 2 {
		log.Fatal("Wrong count of args: host port")
	}

	options.host = flag.Arg(0)
	options.port = flag.Arg(1)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGQUIT)

	go func() {
		sig := <-shutdown
		fmt.Println("Got signal:", sig)
		os.Exit(1)
	}()

	conn, err := connect(options)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := (*conn).Close(); err != nil {
			log.Fatal(err)
		}
	}()

	err = sendRequest(*conn)
	if err != nil {
		log.Fatal(err)
	}
}
