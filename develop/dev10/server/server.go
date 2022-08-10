package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// Сервера подключения
func server(address string, exitChan chan int) {
	l, err := net.Listen("tcp", address)

	if err != nil {
		fmt.Println(err.Error())
		exitChan <- 1
	}

	fmt.Println("listen: " + address)

	defer l.Close()

	for {
		conn, err := l.Accept()

		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		// В соответствии с соединением, чтобы открыть сеанс, этот процесс нуждается в параллельном исполнении.
		go handleSession(conn, exitChan)
	}
}

// обработка разговора
func handleSession(conn net.Conn, exitChan chan int) {
	fmt.Println("Session started")
	reader := bufio.NewReader(conn)

	for {
		str, err := reader.ReadString('\n')

		// Команда Telnet
		if err == nil {
			str = strings.TrimSpace(str)
			if !processTelnetCommand(str, exitChan) {
				conn.Close()
				break
			}

			_, err = conn.Write([]byte(str))
			if err != nil {
				return
			}
		} else {
			// Произошла ошибка
			fmt.Println("Session closed")
			err = conn.Close()
			if err != nil {
				return
			}
			break
		}
	}
}

// Команда протокола телента
func processTelnetCommand(str string, exitChan chan int) bool {
	if strings.HasPrefix(str, "@close") {
		fmt.Println("Session closed")
		// уведомлять внешнюю необходимость отключения
		return false
	} else if strings.HasPrefix(str, "@shutdown") {
		fmt.Println("Server shutdown")
		// Напишите 0 в канал, заблокируйте обработку получателей ожидания
		exitChan <- 0
		return false
	}

	// Распечатать входную строку
	fmt.Println(str)
	return true

}

func main() {
	// Создать канал для конечного кода программы
	exitChan := make(chan int)

	// Одновременно запускайте сервер
	go server("127.0.0.1:8000", exitChan)

	// Блок канала, дождитесь возвращаемого значения
	code := <-exitChan

	// Регистрация программы Возвращаемая стоимость и выходы
	os.Exit(code)
}
