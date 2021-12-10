package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input1")
	if err != nil {
		panic(err)
	}

	one(bufio.NewReader(file))
}

func one(in *bufio.Reader) int {
	for {
		line, err := in.ReadString('\n')
		if err != nil {
			return 0
		}

		fmt.Println(fmt.Sprintf("line: %+v", line))
	}

	return 0
}
