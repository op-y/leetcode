package leetcode

func countPrimes(n int) int {
    if n < 2 {
        return 0
    }

    marks := make([]bool, n)
    for i:=0; i<n; i++ {
        marks = true
    }
    count := 0
    for i:=2; i<n; i++ {
        if marks[i] {
            count++
            if i*i < n {
                for j:=i*i; j<n; j+=i {
                     marks[j] = false
                }
            }
        }
    }
    return count
}

//func isPrime(x int) bool {
//    for i:=2; i*i<=x; i++ {
//        if x % i == 0 {
//            return false
//        }
//    }
//    return true
//}
