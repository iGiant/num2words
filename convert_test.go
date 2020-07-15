package num2words

import (
	"fmt"
	"testing"
)

func TestDivMod(t *testing.T) {
	fmt.Println(divMod(-15))
}

func TestUnitsConvert(t *testing.T) {
	datas := []int{-10, 0, 3, 11, 19, 20, 50, 119, 207, -567, 991, 101}
	var (
		result string
		err    error
	)

	for _, data := range datas {
		result, err = unitsConvert(data, Neuter)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(data, "->", result)
	}
}
