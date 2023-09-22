package leetcode

func intToRoman(num int) string {
	result := ""

	m := 0
	cm := 0
	d := 0
	cd := 0
	c := 0
	xc := 0
	l := 0
	xl := 0
	x := 0
	ix := 0
	v := 0
	iv := 0
	i := 0

	m = num / 1000
	num = num % 1000

	if num >= 900 {
		cm = 1
		num -= 900
	}

	if num >= 500 {
		d = 1
		num -= 500
	}

	if num >= 400 {
		cd = 1
		num -= 400
	}

	c = num / 100
	num = num % 100

	if num >= 90 {
		xc = 1
		num -= 90
	}

	if num >= 50 {
		l = 1
		num -= 50
	}

	if num >= 40 {
		xl = 1
		num -= 40
	}

	x = num / 10
	num = num % 10

	if num >= 9 {
		ix = 1
		num -= 9
	}

	if num >= 5 {
		v = 1
		num -= 5
	}

	if num >= 4 {
		iv = 1
		num -= 4
	}

	i = num

	for idx := 0; idx < m; idx++ {
		result = result + "M"
	}
	if cm == 1 {
		result = result + "CM"
	}
	if d == 1 {
		result = result + "D"
	}
	if cd == 1 {
		result = result + "CD"
	}
	for idx := 0; idx < c; idx++ {
		result = result + "C"
	}
	if xc == 1 {
		result = result + "XC"
	}
	if l == 1 {
		result = result + "L"
	}
	if xl == 1 {
		result = result + "XL"
	}
	for idx := 0; idx < x; idx++ {
		result = result + "X"
	}
	if ix == 1 {
		result = result + "IX"
	}
	if v == 1 {
		result = result + "V"
	}
	if iv == 1 {
		result = result + "IV"
	}
	for idx := 0; idx < i; idx++ {
		result = result + "I"
	}

	return result
}
