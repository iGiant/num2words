package num2words

import (
	"fmt"
)

func unitsConvert(value int, gen gender) (string, error) {
	if value > 999 || value < -999 {
		return "", fmt.Errorf("number limit exceeded (need [-999..999])")
	}
	negative := value < 0
	result := ""
	value, one := divMod(value, 1)
	result = units[one]
	if gen == Feminine && (one == 1 || one == 2) {
		result = feminineUnits[one]
	}
	if gen == Neuter && one == 1 {
		result = neuterUnits[1]
	}
	if value == 0 {
		return insertNegative(result, negative), nil
	}
	if one == 0 {
		result = ""
	}
	value, ten := divMod(value, 1)
	switch ten {
	case 0:
		break
	case 1:
		result = units[ten*10+one]
	default:
		result = tens[ten] + " " + result
	}
	_, hundred := divMod(value, 1)
	if hundred == 0 {
		return insertNegative(result, negative), nil
	}
	return insertNegative(hundreds[hundred]+" "+result, negative), nil
}

func insertNegative(value string, negative bool) string {
	if negative {
		return minus + value
	}
	return value
}

func divMod(value, order int) (int, int) {
	if value == 0 {
		return 0, 0
	}
	pow := 1
	for i := 0; i < order; i++ {
		pow *= 10
	}
	div := value / pow
	mod := value - div*pow
	if mod < 0 {
		mod *= -1
	}
	return div, mod
}

func NumToStr(value int, genderValue gender) (string, error) {
	negative := value < 0
	if negative {
		value *= -1
	}
	value, one := divMod(value, 3)
	if value == 0 {
		result, err := unitsConvert(one, genderValue)
		return insertNegative(result, negative), err
	}
	value, thousand := divMod(value, 3)
	if value == 0 {
		thousandResult, err := unitsConvert(thousand, Feminine)
		if err != nil {
			return "", err
		}
		thousandResult = thousandResult + " " + Enditive(thousand, thousandString[0], thousandString[1], thousandString[2])
		oneResult, err := unitsConvert(one, genderValue)
		if err != nil {
			return "", err
		}
		if oneResult == units[0] {
			oneResult = ""
		}
		return insertNegative(thousandResult+" "+oneResult, negative),
			nil
	}
	return "", nil
}
