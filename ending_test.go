package num2words

import "testing"

func TestEnditive(t *testing.T) {
	const (
		form1 = "a"
		form2 = "b"
		form3 = "c"
	)
	type try struct {
		input  int
		result string
	}
	inputNums := []try{
		try{0, form3},
		try{1, form1},
		try{19, form3},
		try{39, form3},
		try{101, form1},
		try{1011, form3},
		try{1104, form2},
		try{10000000, form3},
	}
	for _, num := range inputNums {
		if Enditive(num.input, form1, form2, form3) != num.result {
			t.Errorf("%d != %s", num.input, num.result)
		}
	}
}
