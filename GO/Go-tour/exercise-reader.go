package main

import "github.com/Go-zh/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(p []byte) (int, error) {
	for i := 0; i < len(p); i++ {
		p[i] = 'A'
	}
	return len(p), nil
}

func main() {
	reader.Validate(MyReader{})
}
