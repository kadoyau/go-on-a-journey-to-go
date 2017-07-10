package main

import "golang.org/x/tour/reader"

// import "fmt"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (MyReader) Read(b []byte) (n int, err error) {
	//	    fmt.Println("%v %v",len(b),cap(b))
	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {

	//	reader := MyReader{}
	//	b := make([]byte, 8)
	//	n, err :=reader.Read(b)
	//	fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
	reader.Validate(MyReader{})
}
