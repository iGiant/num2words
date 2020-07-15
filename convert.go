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
	value, one := divMod(value)
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
	value, ten := divMod(value)
	switch ten {
	case 0:
		break
	case 1:
		result = units[ten*10+one]
	default:
		result = tens[ten] + " " + result
	}
	_, hundred := divMod(value)
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

func divMod(value int) (int, int) {
	if value == 0 {
		return 0, 0
	}
	div := value / 10
	mod := value - div*10
	if mod < 0 {
		mod *= -1
	}
	return div, mod
}
