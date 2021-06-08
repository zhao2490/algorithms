package bitmap

/*
	位图
*/
type BitMap struct {
	bytes []byte // 1byte 占 8位
}

func NewBitMap(length uint) BitMap {
	return BitMap{
		bytes: make([]byte, length/8+1),
	}
}

func (b BitMap) Set(value uint) {
	byteIndex := value / 8
	if byteIndex >= uint(len(b.bytes)) {
		return
	}
	bitIndex := value % 8
	b.bytes[byteIndex] |= 1 << bitIndex
}

func (b BitMap) Get(value uint) bool {
	byteIndex := value / 8
	if byteIndex >= uint(len(b.bytes)) {
		return false
	}
	bitIndex := value % 8
	return b.bytes[byteIndex]&(1<<bitIndex) != 0
}
