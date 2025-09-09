package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}
	linesChan := getLinesChannel(file)

	for line := range linesChan {
		fmt.Println("read:", line)
	}


}

func getLinesChannel(f io.ReadCloser) <-chan string {
	lines := make(chan string)

	go func() {
		defer f.Close()
		defer close(lines)
		currentLinesContents := ""
		for {
			buffer := make([]byte, 8)
			n, err := f.Read(buffer)
			if err != nil {
				if currentLinesContents != "" {
					lines <- currentLinesContents
				}
				if err == io.EOF {
					break
				}
				fmt.Printf("error: %s", err)
				return
			}
			str := string(buffer[:n])
			parts := strings.Split(str, "\n")
			for i := 0; i < len(parts)-1; i++ {
				lines <- fmt.Sprintf("%s%s", currentLinesContents, parts[i])
				currentLinesContents = ""
			}
			currentLinesContents += parts[len(parts)-1]
		}
	}()
	return lines
}
