package protocal

func checkMagic(bytes []byte) bool {
	return bytes[0] == 0xab && bytes[1] == 0xba
}

func copyFullWithOffset(dst []byte, src []byte, start *int) {
	copy(dst[*start:*start+len(src)], src)
	*start = *start + len(src)
}
