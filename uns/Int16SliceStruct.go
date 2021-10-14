package uns

import "unsafe"

var int16Size = int(unsafe.Sizeof(int16(0)))

type Int16SliceStruct SliceStruct

func Int16SliceStructOf(is []int16) *Int16SliceStruct {
	return (*Int16SliceStruct)(unsafe.Pointer(&is))
}

func (iss *Int16SliceStruct) Int16Slice() []int16 {
	return *(*[]int16)(unsafe.Pointer(iss))
}

func (iss *Int16SliceStruct) Bytes() []byte {
	bss := &BytesStruct{Array: iss.Array, Len: iss.Len * int16Size, Cap: iss.Cap * int16Size}
	return *(*[]byte)(unsafe.Pointer(bss))
}

func (iss *Int16SliceStruct) ToBytesStruct() *BytesStruct {
	return &BytesStruct{Array: iss.Array, Len: iss.Len * int16Size, Cap: iss.Cap * int16Size}
}

func Int16SliceToBytes(is []int16) []byte {
	ss := (*SliceStruct)(unsafe.Pointer(&is))
	ss.Len *= int16Size
	ss.Cap *= int16Size
	return *(*[]byte)(unsafe.Pointer(ss))
}
