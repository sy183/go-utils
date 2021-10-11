package uns

import "unsafe"

var int32Size = int(unsafe.Sizeof(int32(0)))

type Int32SliceStruct SliceStruct

func Int32SliceStructOf(is []int32) *Int32SliceStruct {
	return (*Int32SliceStruct)(unsafe.Pointer(&is))
}

func (iss *Int32SliceStruct) Int32Slice() []int32 {
	return *(*[]int32)(unsafe.Pointer(iss))
}

func (iss *Int32SliceStruct) Bytes() []byte {
	bss := &BytesStruct{Array: iss.Array, Len: iss.Len * int32Size, Cap: iss.Cap * int32Size}
	return *(*[]byte)(unsafe.Pointer(bss))
}

func (iss *Int32SliceStruct) ToBytesStruct() *BytesStruct {
	return &BytesStruct{Array: iss.Array, Len: iss.Len * int32Size, Cap: iss.Cap * int32Size}
}

func Int32SliceToBytes(is []int32) []byte {
	ss := (*SliceStruct)(unsafe.Pointer(&is))
	ss.Len *= int32Size
	ss.Cap *= int32Size
	return *(*[]byte)(unsafe.Pointer(ss))
}
