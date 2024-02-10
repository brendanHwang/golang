package page21

import (
	"fmt"
	"strings"
)

func Readers() {
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)

	for {
		n, error := r.Read(b)
		fmt.Printf("n = %v error = %v b = %v\n", n, error, b)
		fmt.Printf("b[:n] = %q\n", b[:n])

		if error != nil {
			break
		}
	}
}
