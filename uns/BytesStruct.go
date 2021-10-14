package uns

import "unsafe"

type BytesStruct SliceStruct

func BytesStructOf(bs []byte) *BytesStruct {
	return (*BytesStruct)(unsafe.Pointer(&bs))
}

func BytesStructOfString(s string) *BytesStruct {
	ss := (*StringStruct)(unsafe.Pointer(&s))
	return &BytesStruct{Array: ss.Str, Len: ss.Len, Cap: ss.Len}
}

func (bss *BytesStruct) Bytes() []byte {
	return *(*[]byte)(unsafe.Pointer(bss))
}

func (bss *BytesStruct) IntSlice() []int {
	iss := &IntSliceStruct{Array: bss.Array, Len: bss.Len / intSize, Cap: bss.Cap / intSize}
	return *(*[]int)(unsafe.Pointer(iss))
}

func (bss *BytesStruct) ToIntSliceStruct() *IntSliceStruct {
	return &IntSliceStruct{Array: bss.Array, Len: bss.Len / intSize, Cap: bss.Cap / intSize}
}

func (bss *BytesStruct) UintSlice() []uint {
	iss := &UintSliceStruct{Array: bss.Array, Len: bss.Len / uintSize, Cap: bss.Cap / uintSize}
	return *(*[]uint)(unsafe.Pointer(iss))
}

func (bss *BytesStruct) ToUintSliceStruct() *UintSliceStruct {
	return &UintSliceStruct{Array: bss.Array, Len: bss.Len / uintSize, Cap: bss.Cap / uintSize}
}

func (bss *BytesStruct) Int8Slice() []int8 {
	iss := &Int8SliceStruct{Array: bss.Array, Len: bss.Len / int8Size, Cap: bss.Cap / int8Size}
	return *(*[]int8)(unsafe.Pointer(iss))
}

func (bss *BytesStruct) ToInt8SliceStruct() *Int8SliceStruct {
	return &Int8SliceStruct{Array: bss.Array, Len: bss.Len / int8Size, Cap: bss.Cap / int8Size}
}

func (bss *BytesStruct) Uint8Slice() []uint8 {
	iss := &Uint8SliceStruct{Array: bss.Array, Len: bss.Len / uint8Size, Cap: bss.Cap / uint8Size}
	return *(*[]uint8)(unsafe.Pointer(iss))
}

func (bss *BytesStruct) ToUint8SliceStruct() *Uint8SliceStruct {
	return &Uint8SliceStruct{Array: bss.Array, Len: bss.Len / uint8Size, Cap: bss.Cap / uint8Size}
}

func (bss *BytesStruct) Int16Slice() []int16 {
	iss := &Int16SliceStruct{Array: bss.Array, Len: bss.Len / int16Size, Cap: bss.Cap / int16Size}
	return *(*[]int16)(unsafe.Pointer(iss))
}

func (bss *BytesStruct) ToInt16SliceStruct() *Int16SliceStruct {
	return &Int16SliceStruct{Array: bss.Array, Len: bss.Len / int16Size, Cap: bss.Cap / int16Size}
}

func (bss *BytesStruct) Uint16Slice() []uint16 {
	iss := &Uint16SliceStruct{Array: bss.Array, Len: bss.Len / uint16Size, Cap: bss.Cap / uint16Size}
	return *(*[]uint16)(unsafe.Pointer(iss))
}

func (bss *BytesStruct) ToUint16SliceStruct() *Uint16SliceStruct {
	return &Uint16SliceStruct{Array: bss.Array, Len: bss.Len / uint16Size, Cap: bss.Cap / uint16Size}
}

func (bss *BytesStruct) Int32Slice() []int32 {
	iss := &Int32SliceStruct{Array: bss.Array, Len: bss.Len / int32Size, Cap: bss.Cap / int32Size}
	return *(*[]int32)(unsafe.Pointer(iss))
}

func (bss *BytesStruct) ToInt32SliceStruct() *Int32SliceStruct {
	return &Int32SliceStruct{Array: bss.Array, Len: bss.Len / int32Size, Cap: bss.Cap / int32Size}
}

func (bss *BytesStruct) Uint32Slice() []uint32 {
	iss := &Uint32SliceStruct{Array: bss.Array, Len: bss.Len / uint32Size, Cap: bss.Cap / uint32Size}
	return *(*[]uint32)(unsafe.Pointer(iss))
}

func (bss *BytesStruct) ToUint32SliceStruct() *Uint32SliceStruct {
	return &Uint32SliceStruct{Array: bss.Array, Len: bss.Len / uint32Size, Cap: bss.Cap / uint32Size}
}

func (bss *BytesStruct) Int64Slice() []int64 {
	iss := &Int64SliceStruct{Array: bss.Array, Len: bss.Len / int64Size, Cap: bss.Cap / int64Size}
	return *(*[]int64)(unsafe.Pointer(iss))
}

func (bss *BytesStruct) ToInt64SliceStruct() *Int64SliceStruct {
	return &Int64SliceStruct{Array: bss.Array, Len: bss.Len / int64Size, Cap: bss.Cap / int64Size}
}

func (bss *BytesStruct) Uint64Slice() []uint64 {
	iss := &Uint64SliceStruct{Array: bss.Array, Len: bss.Len / uint64Size, Cap: bss.Cap / uint64Size}
	return *(*[]uint64)(unsafe.Pointer(iss))
}

func (bss *BytesStruct) ToUint64SliceStruct() *Uint64SliceStruct {
	return &Uint64SliceStruct{Array: bss.Array, Len: bss.Len / uint64Size, Cap: bss.Cap / uint64Size}
}

func (bss *BytesStruct) Float32Slice() []float32 {
	iss := &Float32SliceStruct{Array: bss.Array, Len: bss.Len / float32Size, Cap: bss.Cap / float32Size}
	return *(*[]float32)(unsafe.Pointer(iss))
}

func (bss *BytesStruct) ToFloat32SliceStruct() *Float32SliceStruct {
	return &Float32SliceStruct{Array: bss.Array, Len: bss.Len / float32Size, Cap: bss.Cap / float32Size}
}

func (bss *BytesStruct) Float64Slice() []float64 {
	iss := &Float64SliceStruct{Array: bss.Array, Len: bss.Len / float64Size, Cap: bss.Cap / float64Size}
	return *(*[]float64)(unsafe.Pointer(iss))
}

func (bss *BytesStruct) ToFloat64SliceStruct() *Float64SliceStruct {
	return &Float64SliceStruct{Array: bss.Array, Len: bss.Len / float64Size, Cap: bss.Cap / float64Size}
}

func (bss *BytesStruct) ToString() string {
	return *(*string)(unsafe.Pointer(bss))
}

func (bss *BytesStruct) ToStringStruct() *StringStruct {
	return &StringStruct{Str: bss.Array, Len: bss.Len}
}

func BytesToString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

func BytesToIntSlice(bs []byte) []int {
	ss := (*SliceStruct)(unsafe.Pointer(&bs))
	ss.Len /= intSize
	ss.Cap /= intSize
	return *(*[]int)(unsafe.Pointer(ss))
}

func BytesToUintSlice(bs []byte) []uint {
	ss := (*SliceStruct)(unsafe.Pointer(&bs))
	ss.Len /= uintSize
	ss.Cap /= uintSize
	return *(*[]uint)(unsafe.Pointer(ss))
}

func BytesToInt8Slice(bs []byte) []int8 {
	ss := (*SliceStruct)(unsafe.Pointer(&bs))
	ss.Len /= int8Size
	ss.Cap /= int8Size
	return *(*[]int8)(unsafe.Pointer(ss))
}

func BytesToUint8Slice(bs []byte) []uint8 {
	ss := (*SliceStruct)(unsafe.Pointer(&bs))
	ss.Len /= uint8Size
	ss.Cap /= uint8Size
	return *(*[]uint8)(unsafe.Pointer(ss))
}

func BytesToInt16Slice(bs []byte) []int16 {
	ss := (*SliceStruct)(unsafe.Pointer(&bs))
	ss.Len /= int16Size
	ss.Cap /= int16Size
	return *(*[]int16)(unsafe.Pointer(ss))
}

func BytesToUint16Slice(bs []byte) []uint16 {
	ss := (*SliceStruct)(unsafe.Pointer(&bs))
	ss.Len /= uint16Size
	ss.Cap /= uint16Size
	return *(*[]uint16)(unsafe.Pointer(ss))
}

func BytesToInt32Slice(bs []byte) []int32 {
	ss := (*SliceStruct)(unsafe.Pointer(&bs))
	ss.Len /= int32Size
	ss.Cap /= int32Size
	return *(*[]int32)(unsafe.Pointer(ss))
}

func BytesToUint32Slice(bs []byte) []uint32 {
	ss := (*SliceStruct)(unsafe.Pointer(&bs))
	ss.Len /= uint32Size
	ss.Cap /= uint32Size
	return *(*[]uint32)(unsafe.Pointer(ss))
}

func BytesToInt64Slice(bs []byte) []int64 {
	ss := (*SliceStruct)(unsafe.Pointer(&bs))
	ss.Len /= int64Size
	ss.Cap /= int64Size
	return *(*[]int64)(unsafe.Pointer(ss))
}

func BytesToUint64Slice(bs []byte) []uint64 {
	ss := (*SliceStruct)(unsafe.Pointer(&bs))
	ss.Len /= uint64Size
	ss.Cap /= uint64Size
	return *(*[]uint64)(unsafe.Pointer(ss))
}

func BytesToFloat32Slice(bs []byte) []float32 {
	ss := (*SliceStruct)(unsafe.Pointer(&bs))
	ss.Len /= float32Size
	ss.Cap /= float32Size
	return *(*[]float32)(unsafe.Pointer(ss))
}

func BytesToFloat64Slice(bs []byte) []float64 {
	ss := (*SliceStruct)(unsafe.Pointer(&bs))
	ss.Len /= float64Size
	ss.Cap /= float64Size
	return *(*[]float64)(unsafe.Pointer(ss))
}
