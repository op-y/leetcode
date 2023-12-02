package leetcode

import (
	"fmt"
)

func readBinaryWatch(num int) []string {
	dictHour := map[int][]string{
		0: {"0"},
		1: {"1", "2", "4", "8"},
		2: {"3", "5", "6", "9", "10"},
		3: {"7", "11"}}

	dictMinute := map[int][]string{
		0: {"00"},
		1: {"01", "02", "04", "08", "16", "32"},
		2: {"03", "05", "06", "09", "10", "12", "17", "18", "20", "24", "33", "34", "36", "40", "48"},
		3: {"07", "11", "13", "14", "19", "21", "22", "25", "26", "28", "35", "37", "38", "41", "42", "44", "49", "50", "52", "56"},
		4: {"15", "23", "27", "29", "30", "39", "43", "45", "46", "51", "53", "54", "57", "58"},
		5: {"31", "47", "55", "59"}}

	result := make([]string, 0)

	for i := 0; i <= num; i++ {
		j := num - i
		if hour, okH := dictHour[i]; okH {
			if minute, okM := dictMinute[j]; okM {
				for _, h := range hour {
					for _, m := range minute {
						time := fmt.Sprintf("%s:%s", h, m)
						result = append(result, time)
					}
				}
			}
		}
	}

	return result
}
