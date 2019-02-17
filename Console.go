package console

import (
	"fmt"
)

// Log just take the input and prints them on STDOUT using fmt.Print
func Log(input ...interface{}) {
	for i := 0; i < len(input); i++ {
		fmt.Print(input[i])
	}
	fmt.Print("\n")
}
