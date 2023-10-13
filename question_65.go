package leetcode

func isNumber(s string) bool {
	if !isValid(s) {
		return false
	}
	tokens := split(s)
	return checkTrans(tokens)
}

func checkTrans(tokens []string) bool {
	// 没有任何token
	l := len(tokens)
	if l == 0 {
		return false
	}

	// 不能以E开头
	if tokens[0] == "E" {
		return false
	}

	// 不能以符号和E结尾
	if tokens[l-1] == "S" || tokens[l-1] == "E" {
		return false
	}

	// 检查E部分
	ok, idx := checkE(tokens)
	if !ok {
		return false
	} else {
		if idx > 0 {
			tokens = tokens[0:idx]
		}
	}

	// 检查有效数字部分
	return checkN(tokens)
}

func checkN(tokens []string) bool {
	// 先去掉S, E之前部分后续不能再出现S
	if tokens[0] == "S" {
		tokens = tokens[1:]
	}

	l := len(tokens)
	// 去掉符号位后有效数字部分tokens不能为0或超过3个
	if l == 0 || l > 3 {
		return false
	}

	// 不能再出现S
	for _, t := range tokens {
		if t == "S"  {
			return false
		}
	}

	if l == 1 && tokens[0] != "N" {
		return false
	}
	
	if l == 2 && tokens[0] == tokens[1] {
		return false
	}

	if l == 3 && (tokens[0] != "N" || tokens[1] != "D" || tokens[2] != "N") {
		return false
	}

	return true
}

func checkE(tokens []string) (bool, int) {
	l := len(tokens)
	idx := -1
	count := 0
	for i, item := range tokens {
		if item == "E" {
			count++
			idx = i	
		}
	}
	if count == 0 {
		// 没有e情况
		return true, -1
	}
	if count > 1 {
		// 超过一个e情况
		return false, -1
	}
	// e 后边不可能有超过2个token
	if idx < l-3 {
		return false, -1
	}
	if idx == l-3 && (tokens[l-2] != "S" || tokens[l-1] != "N") {
		return false, -1
	}
	if idx == l-2 && tokens[l-1] != "N" {
		return false, -1
	}
	return true, idx
}

func split(s string) []string {
	tokens := []string{}
	sign := map[rune]bool{
        '-':true,
        '+':true,
	}
	dot := map[rune]bool {
        '.':true,
	}
	e := map[rune]bool {
        'E':true,
        'e':true,
	}
	num := false
	for _, c := range s {
		if _, ok := sign[c]; ok {
			if num {
				tokens = append(tokens, "N")
				num = false
			}
			tokens = append(tokens, "S")
			continue
		}	
		if _, ok := dot[c]; ok {
			if num {
				tokens = append(tokens, "N")
				num = false
			}
			tokens = append(tokens, "D")
			continue
		}	
		if _, ok := e[c]; ok {
			if num {
				tokens = append(tokens, "N")
				num = false
			}
			tokens = append(tokens, "E")
			continue
		}
		num = true
	}
	if num {
		tokens = append(tokens, "N")
		num = false
	}
	return tokens
}


func isValid(s string) bool {
    valid := map[rune]bool{
        '0':true,
        '1':true,
        '2':true,
        '3':true,
        '4':true,
        '5':true,
        '6':true,
        '7':true,
        '8':true,
        '9':true,
        '-':true,
        '+':true,
        '.':true,
        'E':true,
        'e':true,
    }

    for _, c := range s {
        if _, ok := valid[c]; !ok {
            return false
        }
    }
    return true
}
