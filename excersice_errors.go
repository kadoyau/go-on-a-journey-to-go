package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// fmt パッケージは、変数を文字列で出力する際に
	// error インタフェースを確認するのでfloatに変換しないとスタックオーバーフローになる
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	z := 1.0
	var old float64

	for i := 0; (old-z)*(old-z) > 0.00000001; i++ {
		old = z
		z = z - (z*z-x)/2*z
	}

	if x < 0 {
		return z, ErrNegativeSqrt(x)
	} else {
		return z, nil
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
