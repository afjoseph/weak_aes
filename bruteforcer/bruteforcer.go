package bruteforcer

type ActionFunc func([]uint8) bool

func incrementCell(buff []uint8, startIdx uint32, actionFunc ActionFunc) bool {
	var i uint8 = 0x00
	for ; i < 0xff; i++ {
		buff[startIdx] = i
		if !actionFunc(buff) {
			return false
		}
	}
	return true
}

func incrementBuff(buff []uint8, startIdx uint32, actionFunc ActionFunc) bool {
	if startIdx == uint32(len(buff))-1 {
		if !incrementCell(buff, startIdx, actionFunc) {
			return false
		}
	} else {
		var i uint8 = 0x00
		for ; i < 0xff; i++ {
			buff[startIdx] = i
			if !incrementBuff(buff, startIdx+1, actionFunc) {
				return false
			}
		}
	}
	return true
}

func Run(buff []uint8, startIdx uint32, actionFunc ActionFunc) bool {
	for i := startIdx; i < uint32(len(buff)); i++ {
		buff[i] = 0x00
	}
	if !incrementBuff(buff, startIdx, actionFunc) {
		return false
	}
	return true
}
