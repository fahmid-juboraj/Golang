package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	

	elapsed := time.Since(start)
	fmt.Printf("Time elapsed: %s\n", elapsed)
}
