package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type CountingWriterImpl struct {
	w io.Writer
	c int64
}

func (c *CountingWriterImpl) Write (p []byte) (int, error) {
	c.w.Write(p)
	c.c += int64(len(p))
	return len(p), nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := CountingWriterImpl{w: w, c: 0}
	return &cw, &(cw.c)
}

func main() {
	writer, count := CountingWriter(os.Stdout)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			continue
		}
		writer.Write(data)
		fmt.Println(*count)
	}
}
