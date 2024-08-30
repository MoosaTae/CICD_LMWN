package main

import (
	"fmt"

	"github.com/MoosaTae/CICD_LMWN/internal/calculator"
)

func main() {
	sum := calculator.Add(1, 2)
	fmt.Println("Hello, World!", sum)
}
