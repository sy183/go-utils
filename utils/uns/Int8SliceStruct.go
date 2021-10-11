package uns

import "unsafe"

var int8Size = int(unsafe.Sizeof(int8(0)))

type Int8SliceStruct SliceStruct

func Int8SliceStructOf(is []int8) *Int8SliceStruct {
	return (*Int8SliceStruct)(unsafe.Pointer(&is))
}

func (iss *Int8SliceStruct) Int8Slice() []int8 {
	return *(*[]int8)(unsafe.Pointer(iss))
}

func (iss *Int8SliceStruct) Bytes() []byte {
	bss := &BytesStruct{Array: iss.Array, Len: iss.Len * int8Size, Cap: iss.Cap * int8Size}
	return *(*[]byte)(unsafe.Pointer(bss))
}

func (iss *Int8SliceStruct) ToBytesStruct() *BytesStruct {
	return &BytesStruct{Array: iss.Array, Len: iss.Len * int8Size, Cap: iss.Cap * int8Size}
}

func Int8SliceToBytes(is []int8) []byte {
	ss := (*SliceStruct)(unsafe.Pointer(&is))
	ss.Len *= int8Size
	ss.Cap *= int8Size
	return *(*[]byte)(unsafe.Pointer(ss))
}
