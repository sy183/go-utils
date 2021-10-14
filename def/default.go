package def

import "unsafe"

func DefaultStrings(def string, defIf string, strings ...*string) {
	for _, str := range strings {
		if *str == defIf {
			*str = def
		}
	}
}

func DefaultString(def string, defIf string, str *string) {
	if *str == defIf {
		*str = def
	}
}

func DefaultPointer(def unsafe.Pointer, defIf unsafe.Pointer, pointer *unsafe.Pointer) {
	if *pointer == defIf {
		*pointer = def
	}
}

func DefaultPointers(def unsafe.Pointer, defIf unsafe.Pointer, pointers ...*unsafe.Pointer) {
	for _, pointer := range pointers {
		if *pointer == defIf {
			*pointer = def
		}
	}
}
