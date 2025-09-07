package main

import (
	"fmt"
	"runtime"

	"github.com/juliodacer/GoDesdeCero/variables"
)

func main() {
	state, text := variables.ConvertToText(2348)
	fmt.Println(state)
	fmt.Println(text)

	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es Windows, es ", os)
	} else {
		fmt.Println("Esto es Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os)
	}
}
