package uns

import "unsafe"

var int64Size = int(unsafe.Sizeof(int64(0)))

type Int64SliceStruct SliceStruct

func Int64SliceStructOf(is []int64) *Int64SliceStruct {
	return (*Int64SliceStruct)(unsafe.Pointer(&is))
}

func (iss *Int64SliceStruct) Int64Slice() []int64 {
	return *(*[]int64)(unsafe.Pointer(iss))
}

func (iss *Int64SliceStruct) Bytes() []byte {
	bss := &BytesStruct{Array: iss.Array, Len: iss.Len * int64Size, Cap: iss.Cap * int64Size}
	return *(*[]byte)(unsafe.Pointer(bss))
}

func (iss *Int64SliceStruct) ToBytesStruct() *BytesStruct {
	return &BytesStruct{Array: iss.Array, Len: iss.Len * int64Size, Cap: iss.Cap * int64Size}
}

func Int64SliceToBytes(is []int64) []byte {
	ss := (*SliceStruct)(unsafe.Pointer(&is))
	ss.Len *= int64Size
	ss.Cap *= int64Size
	return *(*[]byte)(unsafe.Pointer(ss))
}
