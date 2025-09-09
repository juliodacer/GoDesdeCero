package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func MultiplicationTable() {
	scanner := bufio.NewScanner(os.Stdin)
	var number int
	var err error

	for {
		fmt.Println("Ingrese un numero : ")
		if scanner.Scan() {
			number, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 12; i++ {
		fmt.Printf("%d x %d = %d \n", number, i, i*number)
	}

}
