package simpleerror

import (
	"fmt"
)

func main() {
	oprd1, oprd2 := 10, 0
	result, err := Divide(oprd1, oprd2)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%d / %d = %d\n", oprd1, oprd2, result)
	//OUTPUT: can't divide '10' by zero 10 / 0 = 0
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("can't divide '%d' by zero ", a)
	}
	return a / b, nil
}
