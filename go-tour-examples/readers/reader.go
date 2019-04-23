package main

import (
	"fmt"
	"golang.org/x/tour/reader"
	"io"
	"os"
	"strings"
)

type MyReader struct{}

func (r MyReader) Read(p []byte) (n int, err error) {
	for i := range p {
		p[i] = 'A'
	}

	return len(p), nil
}

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)

	rot32 := func(b byte) byte {
		switch {
		case (b >= 'a' && b <= 'm') || (b >= 'A' && b <= 'M'):
			b = b + 13

		case (b >= 'n' && b <= 'z') || (b >= 'N' && b <= 'Z'):
			b = b - 13
		}
		return b
	}

	for i := range p {
		rot32(p[i])
	}

	return n, err
}

func main() {
	s := "0123456789ABCD"
	r := strings.NewReader(s)

	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Println("==============================================================")
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		fmt.Println("==============================================================")

		if err == io.EOF {
			break
		}
	}

	reader.Validate(MyReader{})

	ss := strings.NewReader("Lbh penpxrq gur pbqr!")
	rr := rot13Reader{ss}
	io.Copy(os.Stdout, &rr)

}
