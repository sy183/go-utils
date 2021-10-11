package uns

import "unsafe"

type SliceStruct struct {
	Array unsafe.Pointer
	Len   int
	Cap   int
}

func SliceStructOfPtr(slicePtr unsafe.Pointer) *SliceStruct {
	return &(*(*SliceStruct)(slicePtr))
}
