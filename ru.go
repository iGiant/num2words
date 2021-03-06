package num2words

type gender int

const (
	Feminine gender = iota
	Masculine
	Neuter
)

var units = [...]string{
	"ноль",
	"один",
	"два",
	"три",
	"четыре",
	"пять",
	"шесть",
	"семь",
	"восемь",
	"девять",
	"десять",
	"одиннадцать",
	"двенадцать",
	"тринадцать",
	"четырнадцать",
	"пятнадцать",
	"шестнадцать",
	"семнадцать",
	"восемнадцать",
	"девятнадцать",
}

var feminineUnits = [...]string{
	"",
	"одна",
	"две",
}

var neuterUnits = [...]string{
	"",
	"одно",
}

var tens = [...]string{
	"",
	"",
	"двадцать",
	"тридцать",
	"сорок",
	"пятьдесят",
	"шестьдесят",
	"семьдесят",
	"восемьдесят",
	"девяносто",
}

var hundreds = [...]string{
	"",
	"сто",
	"двести",
	"триста",
	"четыреста",
	"пятьсот",
	"шестьсот",
	"семьсот",
	"восемьсот",
	"девятьсот",
}

var orderedunitsFeminine = [...]string{
	"нолевая",
	"первая",
	"вторая",
	"третья",
	"четвертая",
	"пятая",
	"шестая",
	"седьмая",
	"восьмая",
	"девятая",
	"десятая",
	"одиннадцатая",
	"двенадцатая",
	"тринадцатая",
	"четырнадцатая",
	"пятнадцатая",
	"шестнадцатая",
	"семнадцатая",
	"восемнадцатая",
	"девятнадцатая",
}

var orderedtensFeminine = [...]string{
	"",
	"",
	"двадцатая",
	"тридцатая",
	"сороковая",
	"пятьдесятая",
	"шестьдесятая",
	"семидесятая",
	"восьмидесятая",
	"девяностая",
}

var orderedhundredsFeminine = [...]string{
	"",
	"сотая",
	"двухсотая",
	"трёхсотая",
	"четырехсотая",
	"пятисотая",
	"шестисотая",
	"семисотая",
	"восьмисотая",
	"девятисотая",
}

var orderedunitsMasculine = [...]string{
	"нолевой",
	"первый",
	"второй",
	"третий",
	"четвертый",
	"пятый",
	"шестой",
	"седьмой",
	"восьмой",
	"девятый",
	"десятый",
	"одиннадцатый",
	"двенадцатый",
	"тринадцатый",
	"четырнадцатый",
	"пятнадцатый",
	"шестнадцатый",
	"семнадцатый",
	"восемнадцатый",
	"девятнадцатый",
}

var orderedtensMasculine = [...]string{
	"",
	"",
	"двадцатый",
	"тридцатый",
	"сороковой",
	"пятидесятый",
	"шестидесятый",
	"семидесятый",
	"восьмидесятый",
	"девяностый",
}

var orderedhundredsMasculine = [...]string{
	"",
	"сотый",
	"двухсотый",
	"трёхсотый",
	"четырехсотый",
	"пятисотый",
	"шестисотый",
	"семисотый",
	"восьмисотый",
	"девятисотый",
}

var orderedunitsNeuter = [...]string{
	"нолевое",
	"первое",
	"второе",
	"третье",
	"четвертое",
	"пятое",
	"шестое",
	"седьмое",
	"восьмое",
	"девятое",
	"десятое",
	"одиннадцатое",
	"двенадцатое",
	"тринадцатое",
	"четырнадцатое",
	"пятнадцатое",
	"шестнадцатое",
	"семнадцатое",
	"восемнадцатое",
	"девятнадцатое",
}

var orderedtensNeuter = [...]string{
	"",
	"",
	"двадцатый",
	"тридцатый",
	"сороковой",
	"пятидесятый",
	"шестидесятый",
	"семидесятый",
	"восьмидесятый",
	"девяностый",
}

var orderedhundredsNeuter = [...]string{
	"",
	"сотое",
	"двухсотое",
	"трёхсотое",
	"четырехсотое",
	"пятисотое",
	"шестисотое",
	"семисотое",
	"восьмисотое",
	"девятисотое",
}

var orderedunitsPlural = [...]string{
	"нолевые",
	"первые",
	"вторые",
	"третьи",
	"четвертые",
	"пятые",
	"шестые",
	"седьмые",
	"восьмые",
	"девятые",
	"десятые",
	"одиннадцатые",
	"двенадцатые",
	"тринадцатые",
	"четырнадцатые",
	"пятнадцатые",
	"шестнадцатые",
	"семнадцатые",
	"восемнадцатые",
	"девятнадцатые",
}

var orderedtensPlural = [...]string{
	"",
	"",
	"двадцатые",
	"тридцатые",
	"сороковые",
	"пятидесятые",
	"шестидесятые",
	"семидесятые",
	"восьмидесятые",
	"девяностые",
}

var orderedhundredsPlural = [...]string{
	"",
	"сотые",
	"двухсотые",
	"трёхсотые",
	"четырехсотые",
	"пятисотые",
	"шестисотые",
	"семисотые",
	"восьмисотые",
	"девятисотые",
}

var unitsPreffix = []string{
	"",
	"одно",
	"двух",
	"трёх",
	"четырёх",
	"пяти",
	"шести",
	"семи",
	"восьми",
	"девяти",
	"десяти",
	"одиннадцати",
	"двенадцати",
	"тринадцати",
	"четырнадцати",
	"пятнадцати",
	"шестнадцати",
	"семнадцати",
	"восемнадцати",
	"девятнадцати",
}

var tensPreffix = [...]string{
	"",
	"",
	"двадцати",
	"тридцати",
	"сорока",
	"пятидесяти",
	"шестидесяти",
	"семидесяти",
	"восьмидесяти",
	"девяносто",
}

var hundredsPreffix = [...]string{
	"",
	"сто",
	"двухсот",
	"трёхсот",
	"четырехсот",
	"пятиста",
	"шестиста",
	"восьмиста",
	"девятиста",
}

var emptyString = [...]string{"", "", ""}

var thousandString = [...]string{" тысяча", " тысячи", " тысяч"}

var millionString = [...]string{" миллион", " миллиона", " миллионов"}

var billionString = [...]string{" миллиард", " миллиарда", " миллиардов"}

var trillionString = [...]string{" триллион", " триллиона", " триллионов"}

var quadrillionString = [...]string{" квадриллион", " квадриллиона", " квадриллионов"}

var quintillionString = [...]string{" квинтиллион", " квинтиллиона", " квинтиллионов"}

var sextillionString = [...]string{" секстиллион", " секстиллиона", " секстиллионов"}

const (
	minus             = "минус "
	thousandFeminine  = "тысячная"
	millionFeminine   = "миллионная"
	billionFeminine   = "миллиардная"
	trillionFeminine  = "триллионная"
	thousandMasculine = "тысячный"
	millionMasculine  = "миллионный"
	billionMasculine  = "миллиардный"
	trillionMasculine = "триллионный"
	thousandNeuter    = "тысячное"
	millionNeuter     = "миллионное"
	billionNeuter     = "миллиардное"
	trillionNeuter    = "триллионное"
)
