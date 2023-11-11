package leetcode

import (
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	res := ""
	if numerator > 0 && denominator < 0 {
		res = "-"
		denominator = -denominator
	}
	if numerator < 0 && denominator > 0 {
		res = "-"
		numerator = -numerator
	}
	if numerator < 0 && denominator < 0 {
		numerator = -numerator
		denominator = -denominator
	}
	quotient := numerator / denominator
	res += strconv.Itoa(quotient)
	remainder := numerator % denominator
	if remainder == 0 {
		return res
	}
	res += "."

	fraction := ""
	m := map[int]int{}
	pos := 0
	for remainder != 0 {
		// 出现重复
		if idx, ok := m[remainder]; ok {
			fraction = fraction[:idx] + "(" + fraction[idx:] + ")"
			break
		}

		m[remainder] = pos
		remainder *= 10
		quotient = remainder / denominator
		remainder = remainder % denominator
		fraction += strconv.Itoa(quotient)
		pos++
	}

	return res + fraction
}
