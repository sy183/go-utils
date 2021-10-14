package uns

import "unsafe"

type StringStruct struct {
	Str unsafe.Pointer
	Len int
}

func StringStructOf(s string) *StringStruct {
	return (*StringStruct)(unsafe.Pointer(&s))
}

func StringStructOfBytes(bs []byte) *StringStruct {
	return (*StringStruct)(unsafe.Pointer(&bs))
}

func (ss *StringStruct) Bytes() []byte {
	return *(*[]byte)(unsafe.Pointer(&SliceStruct{Array: ss.Str, Len: ss.Len, Cap: ss.Len}))
}

func (ss *StringStruct) ToString() string {
	return *(*string)(unsafe.Pointer(ss))
}

func (ss *StringStruct) ToBytesStruct() *BytesStruct {
	return &BytesStruct{Array: ss.Str, Len: ss.Len, Cap: ss.Len}
}

func StringToBytes(s string) []byte {
	ss := (*StringStruct)(unsafe.Pointer(&s))
	return *(*[]byte)(unsafe.Pointer(&SliceStruct{Array: ss.Str, Len: ss.Len, Cap: ss.Len}))
}
