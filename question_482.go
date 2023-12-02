package leetcode

func licenseKeyFormatting(S string, K int) string {
	ss := []byte(S)
	st := make([]byte, 0)
	length := len(ss)
	counter := 0
	for i := length - 1; i >= 0; i-- {
		if S[i] == '-' {
			continue
		} else if S[i] >= 'a' && S[i] <= 'z' {
			c := byte(int(S[i]) - 32)
			st = append([]byte{c}, st...)
			counter++
		} else {
			st = append([]byte{S[i]}, st...)
			counter++
		}
		if counter == K {
			st = append([]byte{'-'}, st...)
			counter = 0
		}
	}
	if len(st) > 0 && st[0] == '-' {
		st = st[1:]
	}
	return string(st)
}
