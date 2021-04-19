package main

import (
	"fmt"

	"github.com/blasvinas/golang/mathops"
)

func main() {
	result := mathops.Add(5, 2)
	fmt.Println(result)
	result = mathops.Subtract(5, 2)
	fmt.Println(result)
}
