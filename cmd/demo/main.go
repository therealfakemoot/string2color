package main

import (
	"flag"
	"fmt"

	"github.com/therealfakemoot/string2color"
)

func main() {
	var input string

	flag.StringVar(&input, "input", "TEST_INPUT", "String to convert to color code.")

	flag.Parse()

	// fmt.Printf("%#+v\n", string2color.ToRGB(input))
	fmt.Printf("%s\n", string2color.ToRGB(input))
	fmt.Printf("%s\n", string2color.ToRGBA(input))
}
