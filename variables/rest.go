package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Name string
var Status bool
var Payment float64
var Date time.Time

func RestVariables() {
	Name = "Pedro"
	Status = true
	Payment = 1577.66
	Date = time.Now()
	fmt.Println(Name)
	fmt.Println(Status)
	fmt.Println(Payment)
	fmt.Println(Date)
}

func ConvertToText(number int) (bool, string) {
	text := strconv.Itoa(number)
	return true, text
}
