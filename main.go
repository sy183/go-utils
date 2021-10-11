package main

import "fmt"

func setFlag(f *[16]byte, n int) {
	f[n >> 3] |= 1 << (n & 7)
}

func main() {
	//is := []int{1, 2, 3, 4, 5, 6, 7, 8}
	//is1 := is[1:3]
	//fmt.Println(is)
	//fmt.Println(is1)
	//fmt.Println(uns.IntSliceStructOf(is))
	//fmt.Println(uns.IntSliceStructOf(is1))
	//is1 = append(is1, 10, 11, 12, 13, 14, 15)
	//fmt.Println(is)
	//fmt.Println(is1)
	//fmt.Println(uns.IntSliceStructOf(is))
	//fmt.Println(uns.IntSliceStructOf(is1))
	f := [16]byte{}
	fmt.Println(f)
	for i := 0; i < 129; i++ {
		setFlag(&f, i)
		fmt.Println(f)
	}
}
