func generateParenthesis(n int) []string {
    left, right := n, n
    return gen(left, right, "")
}

func gen(left, right int, prefix string) []string {
    if left > right {
        return []string{}
    }
    if left == 0 && right == 0 {
        return []string{prefix}
    }
    result := []string{}
    if left > 0 {
        result = append(result, gen(left-1, right, prefix+"(")...)
    }
    if right > 0 {
        result = append(result, gen(left, right-1, prefix+")")...)
    }
    return result
}
