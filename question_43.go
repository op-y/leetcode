package leetcode

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	l1 := len(num1)
	num1Arr := make([]int, l1)
	l2 := len(num2)
	num2Arr := make([]int, l2)
	for i := l1 - 1; i >= 0; i-- {
		n := convert2Num(num1[i])
		num1Arr[l1-(i+1)] = n
	}
	for i := l2 - 1; i >= 0; i-- {
		n := convert2Num(num2[i])
		num2Arr[l2-(i+1)] = n
	}

	lr := l1 + l2 + 1
	result := make([]int, lr)

	for i := 0; i < l1; i++ {
		n1 := num1Arr[i]
		for j := 0; j < l2; j++ {
			n2 := num2Arr[j]
			result[i+j] = result[i+j] + n1*n2
		}
	}

	for i := 0; i < lr; i++ {
		n := result[i]
		if n > 9 {
			result[i] = n % 10
			result[i+1] += n / 10
		}
	}

	for result[len(result)-1] == 0 {
		result = result[0 : len(result)-1]
	}

	s := ""
	lr = len(result)
	for i := 0; i < lr; i++ {
		s = convert2str(result[i]) + s
	}
	return s
}

func convert2Num(c byte) int {
	switch c {
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	default:
		return 0
	}
}

func convert2str(n int) string {
	switch n {
	case 0:
		return "0"
	case 1:
		return "1"
	case 2:
		return "2"
	case 3:
		return "3"
	case 4:
		return "4"
	case 5:
		return "5"
	case 6:
		return "6"
	case 7:
		return "7"
	case 8:
		return "8"
	case 9:
		return "9"
	default:
		return "0"
	}
}

//func main() {
//	num1 := "2"
//	num2 := "3"
//
//	result := multiply(num1, num2)
//	fmt.Printf("result: %s\n", result)
//}
