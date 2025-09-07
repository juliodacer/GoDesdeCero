package variables

import (
	"fmt"
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
