package ejercicios

import (
	"strconv"
)

func ConvNumber(value string) (int, string) {
	newValue, err := strconv.Atoi(value)

	if err != nil {
		return 0, "Hubo un error" + err.Error()
	}

	if newValue > 100 {
		return newValue, "Es mayor a 100"
	} else {
		return newValue, "Es menor a 100"
	}

}
