package num2words

import (
	"math"
	"testing"
)

func TestUnitsConvert(t *testing.T) {
	datas := []int{0, 3, 11, 19, 23, 50, 111119, 202, -567, 991, 101010, 999999, 1000000, 100000001,
		math.MaxInt8, math.MaxInt16, math.MaxInt32, math.MaxInt64, math.MinInt64}
	results := []string{
		"ноль",
		"три",
		"одиннадцать",
		"девятнадцать",
		"двадцать три",
		"пятьдесят",
		"сто одиннадцать тысяч сто девятнадцать",
		"двести два",
		"минус пятьсот шестьдесят семь",
		"девятьсот девяносто один",
		"сто одна тысяча десять",
		"девятьсот девяносто девять тысяч девятьсот девяносто девять",
		"один миллион",
		"сто миллионов один",
		"сто двадцать семь",
		"тридцать две тысячи семьсот шестьдесят семь",
		"два миллиарда сто сорок семь миллионов четыреста восемьдесят три тысячи шестьсот сорок семь",
		"девять квинтиллионов двести двадцать три квадриллиона триста семьдесят два триллиона тридцать шесть миллиардов восемьсот пятьдесят четыре миллиона семьсот семьдесят пять тысяч восемьсот семь",
		"минус девять квинтиллионов двести двадцать три квадриллиона триста семьдесят два триллиона тридцать шесть миллиардов восемьсот пятьдесят четыре миллиона семьсот семьдесят пять тысяч восемьсот восемь",
	}
	var (
		result string
		err    error
	)

	for i, data := range datas {
		result, err = Convert(data, Masculine)
		if err != nil {
			t.Error(err)
		}
		if result != results[i] {
			t.Errorf("%d != %s\n", data, result)
		}
	}
}
