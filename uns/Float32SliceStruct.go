package uns

import "unsafe"

var float32Size = int(unsafe.Sizeof(float32(0)))

type Float32SliceStruct SliceStruct

func Float32SliceStructOf(fs []float32) *Float32SliceStruct {
	return (*Float32SliceStruct)(unsafe.Pointer(&fs))
}

func (fss *Float32SliceStruct) Float32Slice() []float32 {
	return *(*[]float32)(unsafe.Pointer(fss))
}

func (fss *Float32SliceStruct) Bytes() []byte {
	bss := &BytesStruct{Array: fss.Array, Len: fss.Len * float32Size, Cap: fss.Cap * float32Size}
	return *(*[]byte)(unsafe.Pointer(bss))
}

func (fss *Float32SliceStruct) ToBytesStruct() *BytesStruct {
	return &BytesStruct{Array: fss.Array, Len: fss.Len * float32Size, Cap: fss.Cap * float32Size}
}

func Float32SliceToBytes(fs []float32) []byte {
	ss := (*SliceStruct)(unsafe.Pointer(&fs))
	ss.Len *= float32Size
	ss.Cap *= float32Size
	return *(*[]byte)(unsafe.Pointer(ss))
}
