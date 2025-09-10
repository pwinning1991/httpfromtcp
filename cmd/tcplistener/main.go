package main

import (
	"fmt"
	request "github.com/pwinning1991/httpfromtcp/internal/request"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		requestLine, err := request.RequestFromReader(conn)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Request line:\n- Method %s\n- Target: %s\n- Version: %s\n", requestLine.RequestLine.Method, requestLine.RequestLine.RequestTarget, requestLine.RequestLine.HttpVersion)
	}

}

//func getLinesChannel(f io.ReadCloser) <-chan string {
//	lines := make(chan string)
//
//	go func() {
//		defer close(lines)
//		defer f.Close()
//		currentLinesContents := ""
//		for {
//			buffer := make([]byte, 8)
//			n, err := f.Read(buffer)
//			if err != nil {
//				if currentLinesContents != "" {
//					lines <- currentLinesContents
//				}
//				if err == io.EOF {
//					break
//				}
//				fmt.Printf("error: %s", err)
//				return
//			}
//			str := string(buffer[:n])
//			parts := strings.Split(str, "\n")
//			for i := 0; i < len(parts)-1; i++ {
//				lines <- fmt.Sprintf("%s%s", currentLinesContents, parts[i])
//				currentLinesContents = ""
//			}
//			currentLinesContents += parts[len(parts)-1]
//		}
//	}()
//	return lines
//}
