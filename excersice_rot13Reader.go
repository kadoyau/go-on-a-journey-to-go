package main

import (
	"io"
	"os"
	"strings"
	//	"fmt"
	"bytes"
)

var ascii_uppercase = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var ascii_lowercase = []byte("abcdefghijklmnopqrstuvwxyz")
var alphabet_length = len(ascii_uppercase)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	pos := bytes.IndexByte(ascii_uppercase, b)
	if pos != -1 {
		return ascii_uppercase[(pos+13)%alphabet_length]
	}

	pos = bytes.IndexByte(ascii_lowercase, b)
	if pos != -1 {
		return ascii_uppercase[(pos+13)%alphabet_length]
	}

	return b
}

```
@see https://gist.github.com/flc/6439105
rot13Reader変数rをReadしたときここが呼ばれる
```
func (r rot13Reader) Read(p []byte) (n int, err error) {
	// NewReaderで生成した*Reader r.r（s） のデータを 
  // *strings.Reader.Read()をつかって p に入れる
	// https://golang.org/pkg/strings/#Reader.Read
	n, err = r.r.Read(p) 

	// convert given the string
	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!") // returns *strings.Reader
	r := rot13Reader{s} // r.r is io.Reader(*strings.Reader)
	io.Copy(os.Stdout, &r)
}
