package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// Do some work here...

	elapsed := time.Since(start)
	fmt.Printf("Time elapsed: %s\n", elapsed)
}
