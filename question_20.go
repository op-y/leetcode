package leetcode

func isValid(s string) bool {
    if len(s) == 0 {
        return true
    }
    stack := make([]byte, 0)
    for _, c := range []byte(s) {
        if c == '{' || c == '[' || c == '(' {
            stack = append(stack, c)
        }
        if c == '}' {
            if len(stack) == 0 {
                return false
            }
            top := stack[len(stack) - 1]
            if top != '{' {
                return false
            }
            stack = stack[0:len(stack) - 1]
        }
        if c == ']' {
            if len(stack) == 0 {
                return false
            }
            top := stack[len(stack) - 1]
            if top != '[' {
                return false
            }
            stack = stack[0:len(stack) - 1]
        }
        if c == ')' {
            if len(stack) == 0 {
                return false
            }
            top := stack[len(stack) - 1]
            if top != '(' {
                return false
            }
            stack = stack[0:len(stack) - 1]
        }
    }
    if len(stack) > 0 {
        return false
    }
    return true
}
