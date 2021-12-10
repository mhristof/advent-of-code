package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/stretchr/testify/assert"
)

func openFile(t *testing.T, name string) *bufio.Reader {
	file, err := os.Open(name)
	if err != nil {
		t.Fatal(err)
	}
	return bufio.NewReader(file)
}

func TestOne(t *testing.T) {
	cases := []struct {
		name string
		in   *bufio.Reader
		out  int
	}{
		{
			name: "example",
			in: bufio.NewReader(strings.NewReader(heredoc.Doc(`
				199
				200
				208
				210
				200
				207
				240
				269
				260
				263
			`))),
			out: 7,
		},
		{
			name: "no increment",
			in: bufio.NewReader(strings.NewReader(heredoc.Doc(`
				269
				263
				260
				240
				207
				200
			`))),
			out: 0,
		},
		{
			name: "all increase",
			in: bufio.NewReader(strings.NewReader(heredoc.Doc(`
				200
				207
				240
				260
				263
				269
			`))),
			out: 5,
		},
		{
			name: "input",
			in:   openFile(t, "input1"),
			out:  1226,
		},
	}

	for _, test := range cases {
		assert.Equal(t, test.out, one(test.in), test.name)
	}
}
