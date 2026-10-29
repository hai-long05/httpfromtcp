package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	for {
		buf := make([]byte, 8)
		n, err := f.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Printf("read: %s\n", string(buf[:n]))
	}
}
