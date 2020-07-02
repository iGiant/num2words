package num2words

import (
	"fmt"
)

func unitsConvert(value int) (string, error) {
	if value > 1000 || value < -1000 {
		return "", fmt.Errorf("number limit exceeded (need [-999..999])")
	}
	result := ""
	value, one := divMod(value)
	result = units[one]
	if value == 0 {
		return result, nil
	}
	if one == 0 {
		result = ""
	}
	value, ten := divMod(value)
	switch ten {
	case 0: result
	}
		result
	}
	result = tens[ten] + " " + result
	return result, nil
}

func divMod(value int) (int, int) {
	if value == 0 {
		return 0, 0
	}
	div := value / 10
	mod := value - div * 10
	if mod < 0 {
		mod *= -1
	}
	return div, mod
}