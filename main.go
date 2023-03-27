package main

import (
	"fmt"

	"github.com/ariszzz/hello/test"
)

func main() {
	i := 33

	sum := test.Add(5, i)
	fmt.Printf("5 + %d = %d\n", i, sum)
	test.Dele()

}
