package num2words

// Enditive возвращает правильную форму существительного в зависимости от числа num
// form1 - соответствует 1 шт, form2 - соответствует от 2 до 4 шт, form3 - остальные
// например Enditive(119, "яблоко", "яблока", "яблок") вернет "яблок"
func Enditive(num int, form1, form2, form3 string) string {
	if num < 0 {
		num *= -1
	}
	mod := num % 100
	if mod >= 11 && mod <= 20 {
		return form3
	}
	mod = num % 10
	if mod == 0 || mod >= 5 {
		return form3
	}
	if mod == 1 {
		return form1
	}
	return form2
}
