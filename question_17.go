func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }
    result := []string{}
    for i:=0; i<len(digits); i++ {
        result = moreDigit(digits[i], result)
    }
    return result
}

func moreDigit(c byte, s []string) []string {
    result := []string{}
    switch c {
        case '2':
            if len(s) == 0 {
                result = []string{"a", "b", "c"}
            } else {
                for _, item := range s {
                    result = append(result, item+"a", item+"b", item+"c")
                }
            }
        case '3':
            if len(s) == 0 {
                result = []string{"d", "e", "f"}
            } else {
                for _, item := range s {
                    result = append(result, item+"d", item+"e", item+"f")
                }
            }
        case '4':
            if len(s) == 0 {
                result = []string{"g", "h", "i"}
            } else {
                for _, item := range s {
                    result = append(result, item+"g", item+"h", item+"i")
                }
            }
        case '5':
            if len(s) == 0 {
                result = []string{"j", "k", "l"}
            } else {
                for _, item := range s {
                    result = append(result, item+"j", item+"k", item+"l")
                }
            }
        case '6':
            if len(s) == 0 {
                result = []string{"m", "n", "o"}
            } else {
                for _, item := range s {
                    result = append(result, item+"m", item+"n", item+"o")
                }
            }
        case '7':
            if len(s) == 0 {
                result = []string{"p", "q", "r", "s"}
            } else {
                for _, item := range s {
                    result = append(result, item+"p", item+"q", item+"r", item+"s")
                }
            }
        case '8':
            if len(s) == 0 {
                result = []string{"t", "u", "v"}
            } else {
                for _, item := range s {
                    result = append(result, item+"t", item+"u", item+"v")
                }
            }
        case '9':
            if len(s) == 0 {
                result = []string{"w", "x", "y", "z"}
            } else {
                for _, item := range s {
                    result = append(result, item+"w", item+"x", item+"y", item+"z")
                }
            }
    }
    return result
}
