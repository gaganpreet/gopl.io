package main

import (
	"bufio"
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
	scanner := bufio.NewScanner(c)
	cwd := "."
	for scanner.Scan() {
		text := scanner.Text()
		if text == "close" {
			break
		} else if text == "ls" {
			fileList := getFileList(cwd)
			for _, filename := range fileList {
				io.WriteString(c, filename+"\n")
			}
		} else if text[:3] == "get" {
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
		} else if text[:2] == "cd" {
			res := strings.Split(text, " ")
			if len(res) != 2 {
				io.WriteString(c, "Invalid input for cd")
			}
			newPath := res[1]
			cwd = path.Join(cwd, newPath)
		} else {
			io.WriteString(c, "Invalid command")
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
