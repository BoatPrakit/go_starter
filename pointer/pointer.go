package pointer

import (
	"fmt"
)

func Pointer(x int) {
	i := 1
	ChangeValue(&i)
	fmt.Println(i)
}

func ChangeValue(a *int) {
	*a = 2
}
