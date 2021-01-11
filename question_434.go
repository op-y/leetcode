package leetcode

func countSegments(s string) int {
     if s == "" {
         return 0
     }
     sum := 0
     isLastBlank := true
     isThisBlank := false
     ss := []byte(s)
     for _, b := range(ss) {
         if b == ' ' {
             isThisBlank = true
         } else {
             isThisBlank = false
         }
         if isLastBlank && !isThisBlank {
             sum += 1
         }
         isLastBlank = isThisBlank
     }
     return sum
}
