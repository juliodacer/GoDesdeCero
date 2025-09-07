package main

import (
	"fmt"

	"github.com/juliodacer/GoDesdeCero/variables"
)

func main() {
	state, text := variables.ConvertToText(2348)
	fmt.Println(state)
	fmt.Println(text)
}
