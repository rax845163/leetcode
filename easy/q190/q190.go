package q190

func reverseBits(num uint32) uint32 {
	ret := uint32(0)
	for i := uint8(1); i <= 32; i++ {
		ret = (ret << 1) | (num & 1)
		num = num >> 1
	}
	return ret
}
