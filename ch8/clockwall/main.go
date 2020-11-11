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

/*
* TZ=US/Eastern gor ch8/clock1/main.go -port 8010&
* TZ=Asia/Tokyo gor ch8/clock1/main.go -port 8020&
* TZ=Europe/London gor ch8/clock1/main.go -port 8030&
 */

func main() {
	for index, arg := range os.Args[1:] {
		res := strings.Split(arg, "=")
		city := res[0]
		fmt.Printf("% *s\n", index*20, city)
	}
	for index, arg := range os.Args[1:] {
		res := strings.Split(arg, "=")
		city, address := res[0], res[1]
		go connectToClock(city, address, index)
	}
	fmt.Printf("\n")
	for {
		// Loop forever
		time.Sleep(1)
	}
}

func connectToClock(city, address string, index int) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Printf("% *s\n", index*20, scanner.Text())
	}

}
