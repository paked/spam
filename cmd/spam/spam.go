package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/paked/spam"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	s := spam.New(os.Stdin)

	fmt.Print(s)
}
