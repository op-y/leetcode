package leetcode

func twoSum(numbers []int, target int) []int {
    if numbers == nil {
        return nil
    }

    if len(numbers) < 2 {
        return nil
    }

    indexA := 0
    indexB := len(numbers) - 1

    for indexA < indexB {
        if numbers[indexA] + numbers[indexB] < target {
            indexA++
        } else if numbers[indexA] + numbers[indexB] > target {
            indexB--
        } else {
            return []int{indexA+1, indexB+1}
        }
    }
    return nil
}
