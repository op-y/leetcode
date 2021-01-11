package leetcode

func reverseBits(num uint32) uint32 {
    data := uint32(0)
    for i:=0;i<16;i++ {
        bit := ( (uint32(1) << i) & num ) << (31-2*i)
        data |= bit
    }
    for i:=16;i<32;i++ {
        bit := ( (uint32(1) << i) & num ) >> (2*i-31)
        data |= bit
    }
    return data
}
