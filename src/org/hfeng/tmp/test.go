package main

import (
	"fmt"
)

func main() {
	sum := 15
	switch {
	case sum > 10:
		fmt.Println("OK")
	case sum > 12:
		fmt.Println("Rather OK")
	}
}
