package uns

import "unsafe"

var float64Size = int(unsafe.Sizeof(float64(0)))

type Float64SliceStruct SliceStruct

func Float64SliceStructOf(fs []float64) *Float64SliceStruct {
	return (*Float64SliceStruct)(unsafe.Pointer(&fs))
}

func (fss *Float64SliceStruct) Float64Slice() []float64 {
	return *(*[]float64)(unsafe.Pointer(fss))
}

func (fss *Float64SliceStruct) Bytes() []byte {
	bss := &BytesStruct{Array: fss.Array, Len: fss.Len * float64Size, Cap: fss.Cap * float64Size}
	return *(*[]byte)(unsafe.Pointer(bss))
}

func (fss *Float64SliceStruct) ToBytesStruct() *BytesStruct {
	return &BytesStruct{Array: fss.Array, Len: fss.Len * float64Size, Cap: fss.Cap * float64Size}
}

func Float64SliceToBytes(fs []float64) []byte {
	ss := (*SliceStruct)(unsafe.Pointer(&fs))
	ss.Len *= float64Size
	ss.Cap *= float64Size
	return *(*[]byte)(unsafe.Pointer(ss))
}
