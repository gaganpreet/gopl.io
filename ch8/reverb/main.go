package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8002")
	if err != nil {
		log.Fatal(os.Stderr, "err: %s", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(os.Stderr, "err: %s", err)
		}
		go handleConn(conn)
	}

}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	// Note: Ignoring potential errors from input.Err()
	c.Close()
}

func echo(c net.Conn, text string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(text))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", text)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(text))
}
