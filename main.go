package main

import (
	"bufio"
	"io"
	"math"
	"strconv"
	"strings"
)

func main() {
}

func one(in *bufio.Reader) (ret int) {
	old := math.MaxUint32

	for {
		line, err := in.ReadString('\n')
		if line == "" || (err != nil && err != io.EOF) {
			return ret
		}
		eof := err == io.EOF

		number, err := strconv.Atoi(strings.TrimSuffix(line, "\n"))
		if err != nil {
			panic(err)
		}

		if number > old {
			ret++
		}

		if eof {
			return ret
		}

		//fmt.Println("old:", old, "number", number, "ret", ret)

		old = number
	}
}
