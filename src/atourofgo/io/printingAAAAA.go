package main

import "code.google.com/p/go-tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (m MyReader) Read(b []byte) (int, error) {
	for i, _ := range b {
		b[i] = 'A'
	}
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}