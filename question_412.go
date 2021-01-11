package leetcode

import (
    "fmt"
)

func fizzBuzz(n int) []string {
    if n <= 0 {
        return nil
    }
    result := make([]string, 0)
    for i:=1;i<=n;i++ {
        if i % 15 == 0 {
            result = append(result, "FizzBuzz")
        } else if i % 5 == 0 {
            result = append(result, "Buzz")
        } else if i % 3 == 0 {
            result = append(result, "Fizz")
        } else {
            result = append(result, fmt.Sprintf("%d", i))
        }
    }
    return result
}
