package main

import (
	"fmt"
	"os"

	"github.com/paked/spam"
)

func main() {
	s := spam.New(os.Stdin)

	fmt.Print(s)
}
