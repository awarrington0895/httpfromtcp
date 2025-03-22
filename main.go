package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
)

const inputFilePath = "messages.txt"

func getLinesChannel(f io.ReadCloser) <-chan string {
	lines := make(chan string)
	go func() {
		currentLineContents := ""
		for {
			buffer := make([]byte, 8, 8)
			n, err := f.Read(buffer)
			if err != nil {
				if currentLineContents != "" {
					lines <- currentLineContents
					currentLineContents = ""
				}
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("error: %s\n", err.Error())
				break
			}
			str := string(buffer[:n])
			parts := strings.Split(str, "\n")
			for i := 0; i < len(parts)-1; i++ {
				lines <- fmt.Sprintf("%s%s", currentLineContents, parts[i])
				currentLineContents = ""
			}
			currentLineContents += parts[len(parts)-1]
		}

		close(lines)
		f.Close()
	}()

	return lines
}

func main() {
	listener, err := net.Listen("tcp", "localhost:42069")
	defer listener.Close()

	if err != nil {
		fmt.Println("Unable to open listener!")
		panic(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("unable to accept connection")
			panic(err)
		}

		fmt.Println("connection successfully established")
		connChan := getLinesChannel(conn)
		for line := range connChan {
			fmt.Println(line)
		}

		fmt.Println("Connection closed!")
	}

	// lineChannel := getLinesChannel(f)
	// for line := range lineChannel {
	// 	fmt.Printf("read: %s\n", line)
	// }
}
