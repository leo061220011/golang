package main

import (
	"fmt"
	"math"
)

func main() {

	a := 260

	if a >= math.MinInt8 && a <= math.MaxInt8 {
		fmt.Printf("true")

	} else {
		fmt.Printf("false")
	}

}
