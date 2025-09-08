package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}
	for {
		words := make([]byte, 8)
		_, err := file.Read(words)
		if err == io.EOF {
			break
		}
		s := fmt.Sprintf("read: %s\n", words)
		io.WriteString(os.Stdout, s)
	}
}

