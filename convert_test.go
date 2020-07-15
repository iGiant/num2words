package num2words

import (
	"fmt"
	"testing"
)

func TestDivMod(t *testing.T) {
	fmt.Println(divMod(-15, 1))
}

func TestUnitsConvert(t *testing.T) {
	datas := []int{-10, 0, 3, 11, 19, 20, 50, 111119, 202, -567, 991, 101000, 999999}
	var (
		result string
		err    error
	)

	for _, data := range datas {
		result, err = NumToStr(data, Feminine)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(data, "->", result)
	}
}
