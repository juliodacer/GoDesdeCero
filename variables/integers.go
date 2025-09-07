package variables

import "fmt"

// Para que pueda ser llamada desde fuera tiene que estar en mayúscula
func ShowIntegers() {
	// manera declarativa
	var commonInt int
	// por asignación
	intOf32 := int32(10)
	intOf64 := int64(100)

	fmt.Println("commonInt = ", commonInt)
	fmt.Println("intOf32 = ", intOf32)
	fmt.Println("intOf64 = ", intOf64)
}
