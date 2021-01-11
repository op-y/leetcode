package

func getSum(a int, b int) int {
    sum := a ^ b
    carrier := (a & b) << 1
    for carrier != 0 {
       tmp := sum ^ carrier
       carrier = (sum & carrier) << 1
       sum = tmp
    }
    return sum
}
