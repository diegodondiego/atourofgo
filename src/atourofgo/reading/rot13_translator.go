package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

type SimpleError string

func (err SimpleError) Error() string {
	return fmt.Sprintf("Error ocurred: %f", err)
}

func (rot13 *rot13Reader) Read(p []byte) (n int, err error) {

	n, err = rot13.r.Read(p)

	if p == nil {
		return 0, SimpleError("array is nil")
	}

	for i := range p {
		if p[i] >= 'A' && p[i] <= 'z' {
			// 'N' the letter's index is between A - N, add 13 to its index
			if (p[i] < 'N') || (p[i] >= 'a' && p[i] < 'n') {
				p[i] += 13

			// if the letter's index is between M - Z, subtract 13 from its index
			} else if (p[i] > 'M' && p[i] <= 'Z') || (p[i] > 'm') {
				p[i] -= 13
			}
		}
	}

	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
