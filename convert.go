package num2words

import (
	"fmt"
	"strings"
)

func unitsConvert(value int, gen gender) (string, error) {
	if value > 999 {
		return "", fmt.Errorf("number limit exceeded (need < 1000)")
	}
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
		return result, nil
	}
	if one == 0 {
		result = ""
	} else {
		result = " " + result
	}
	value, ten := divMod(value, 1)
	switch ten {
	case 0:
		break
	case 1:
		result = units[ten*10+one]
	default:
		result = tens[ten] + result
	}
	_, hundred := divMod(value, 1)
	if hundred == 0 {
		return result, nil
	}
	if result != "" {
		result = " " + strings.TrimSpace(result)
	}
	return hundreds[hundred] + result, nil
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

func addOrder(intValue int, strValue string, order [3]string) string {
	if intValue == 0 {
		return ""
	}
	return strValue + Enditive(intValue, order[0], order[1], order[2])
}

func Convert(value int, genderValue gender) (string, error) {
	negative := value < 0
	if negative {
		value *= -1
	}
	result := make([]string, 0)
	value, one := divMod(value, 3)
	oneResult, err := unitsConvert(one, genderValue)
	if err != nil {
		return "", err
	}
	oneResult = addOrder(one, oneResult, emptyString)
	if oneResult != "" {
		result = append(result, oneResult)
	}
	if value == 0 {
		if oneResult == "" {
			oneResult = units[0]
		}
		return insertNegative(oneResult, negative), err
	}
	value, err = getCapacity(&result, value, thousandString, Feminine)
	if value == 0 {
		return insertNegative(strings.Join(result, " "), negative), nil
	}
	value, err = getCapacity(&result, value, millionString, Masculine)
	if value == 0 {
		return insertNegative(strings.Join(result, " "), negative), nil
	}
	value, err = getCapacity(&result, value, billionString, Masculine)
	if value == 0 {
		return insertNegative(strings.Join(result, " "), negative), nil
	}
	value, err = getCapacity(&result, value, trillionString, Masculine)
	if value == 0 {
		return insertNegative(strings.Join(result, " "), negative), nil
	}
	value, err = getCapacity(&result, value, quadrillionString, Masculine)
	if value == 0 {
		return insertNegative(strings.Join(result, " "), negative), nil
	}
	value, err = getCapacity(&result, value, quintillionString, Masculine)
	if value == 0 {
		return insertNegative(strings.Join(result, " "), negative), nil
	}
	return "", fmt.Errorf("number is big")
}

func getCapacity(result *[]string, value int, capacityName [3]string, gen gender) (int, error) {
	value, capacity := divMod(value, 3)
	capacityResult, err := unitsConvert(capacity, gen)
	if err != nil {
		return 0, err
	}
	capacityResult = addOrder(capacity, capacityResult, capacityName)
	if capacityResult != "" {
		*result = append([]string{capacityResult}, *result...)
	}
	return value, nil
}
