package main

import (
	"bufio"
	"log"
	"math/rand"
	"net"
	"strings"
	"time"
)

// telnet localhost port for check
// сетевой адрес и протокол
const addr = "0.0.0.0:12345"
const proto = "tcp4"

func main() {
	// поговорки для отправки клиенту
	var proverbs []string = []string{"Don't communicate by sharing memory, share memory by communicating.", "Concurrency is not parallelism.", "Channels orchestrate; mutexes serialize.", "The bigger the interface, the weaker the abstraction.", "Make the zero value useful.", "Interface{} says nothing.", "Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.", "A little copying is better than a little dependency.", "Syscall must always be guarded with build tags.", "Cgo must always be guarded with build tags.", "Cgo is not Go.", "With the unsafe package there are no guarantees.", "Clear is better than clever.", "Reflection is never clear.", "Errors are values.", "Don't just check errors, handle them gracefully.", "Design the architecture, name the components, document the details.", "Documentation is for users.", "Don't panic."}
	// запуск сетевой службы
	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	// обработка подключений
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn, proverbs)
	}
}

// функция обработки подключений
func handleConn(conn net.Conn, proverbs []string) {
	for {
		// горутина ждет сообщения для завершения соединения
		go exit(conn)
		// раз в 3 секунды посылает рандомную поговорку
		time.Sleep(3 * time.Second)
		randV := rand.Intn(len(proverbs))
		conn.Write([]byte("> " + proverbs[randV] + "\n"))
	}
}

// функция для завершения соединения
func exit(conn net.Conn) {
	reader := bufio.NewReader(conn)
	b, err := reader.ReadBytes('\n')
	if err != nil {
		log.Println(err)
		return
	}
	msg := strings.TrimSuffix(string(b), "\n")
	msg = strings.TrimSuffix(msg, "\r")

	switch {
	case msg == "quit":
		conn.Close()
	default:
		conn.Write([]byte("> type quit for terminate connection\n"))
	}
}
