package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:2222")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	io.WriteString(c, "220 Hello\n")
	scanner := bufio.NewScanner(c)
	cwd := "."
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("Received ", text)
		command := strings.SplitN(text, " ", 2)[0]
		command = strings.ToLower(command)
		switch command {
		case "close":
			break
		case "ls":
			fileList := getFileList(cwd)
			for _, filename := range fileList {
				io.WriteString(c, filename+"\n")
			}
		case "get":
			res := strings.Split(text, " ")
			if len(res) != 2 {
				io.WriteString(c, "Invalid input for get")
			}
			filename := res[1]
			file, err := os.Open(path.Join(cwd, filename))
			if err != nil {
				io.WriteString(c, "File not found")
			}
			io.Copy(c, file)
		case "cd":
			res := strings.Split(text, " ")
			if len(res) != 2 {
				io.WriteString(c, "Invalid input for cd")
			}
			newPath := res[1]
			cwd = path.Join(cwd, newPath)
		case "user":
			io.WriteString(c, "331 Anonymous login ok")
		case "pass":
			io.WriteString(c, "230 Anonymous access granted")
		case "pwd":
			io.WriteString(c, "250 CWD command successful")
		case "epsv":
		case "pasv":
			io.WriteString(c, "227 Entering Passive Mode")
		case "syst":
			io.WriteString(c, "215 UNIX Type: L8")
		default:
			io.WriteString(c, "502 Command not implemented")
			fmt.Fprintf(os.Stderr, "Invalid command: %s\n", command)
		}
		io.WriteString(c, "\n")
	}
}

func getFileList(cwd string) []string {
	var fileList []string
	files, err := ioutil.ReadDir(cwd)
	if err != nil {
		return fileList
	}
	for _, file := range files {
		name := file.Name()
		if file.IsDir() {
			name = name + "/"
		}
		fileList = append(fileList, name)
	}
	return fileList
}
