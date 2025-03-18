package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./messages.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	buf := make([]byte, 8)

	for {
		_, err = file.Read(buf)

		if err != nil {
			if errors.Is(err, io.EOF) {
				os.Exit(0)
			}

			panic(err)
		}

		fmt.Printf("read: %s\n", string(buf))
	}

}
