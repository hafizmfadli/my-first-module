package main

import (
	"fmt"

	"github.com/hafizmfadli/my-first-module/formatter"
	"github.com/hafizmfadli/my-first-module/math"
)

func main() {
	var num int
	fmt.Scan(&num)
	result := formatter.Format(math.Cube(num))
	fmt.Println(result)
}
