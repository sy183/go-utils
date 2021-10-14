package uns

import "unsafe"

var intSize = int(unsafe.Sizeof(0))

type IntSliceStruct SliceStruct

func IntSliceStructOf(is []int) *IntSliceStruct {
	return (*IntSliceStruct)(unsafe.Pointer(&is))
}

func (iss *IntSliceStruct) IntSlice() []int {
	return *(*[]int)(unsafe.Pointer(iss))
}

func (iss *IntSliceStruct) Bytes() []byte {
	bss := &BytesStruct{Array: iss.Array, Len: iss.Len * intSize, Cap: iss.Cap * intSize}
	return *(*[]byte)(unsafe.Pointer(bss))
}

func (iss *IntSliceStruct) ToBytesStruct() *BytesStruct {
	return &BytesStruct{Array: iss.Array, Len: iss.Len * intSize, Cap: iss.Cap * intSize}
}

func IntSliceToBytes(is []int) []byte {
	ss := (*SliceStruct)(unsafe.Pointer(&is))
	ss.Len *= intSize
	ss.Cap *= intSize
	return *(*[]byte)(unsafe.Pointer(ss))
}
