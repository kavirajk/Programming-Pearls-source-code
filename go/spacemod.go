package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

func measure[T any](name string) {
	fmt.Printf("%s\t%d\t", name, unsafe.Sizeof(*new(T)))

	const N = 11
	ptrs := make([]uintptr, 0, N)
	keep := make([]*T, 0, N)

	for _ = range N {
		p := new(T)
		thisp := uintptr(unsafe.Pointer(p))
		ptrs = append(ptrs, thisp)
		keep = append(keep, p)
		runtime.KeepAlive(p)
	}

	for i := 1; i < len(ptrs); i++ {
		fmt.Printf(" %d", ptrs[i]-ptrs[i-1])
	}

	fmt.Println()
}

type structb struct{ b byte }
type structib struct {
	i int
	b byte
}
type structip struct {
	i int
	p *structip
}
type structi32b struct {
	i int32
	b byte
}
type structi16b struct {
	i int16
	b byte
}
type structbib struct {
	b1 byte
	i  int
	b2 byte
}
type structiii struct {
	i1 int
	i2 int
	i3 int
}
type structiib struct {
	i1 int
	i2 int
	b  byte
}
type structb12 struct {
	c [12]byte
}
type structb13 struct {
	c [13]byte
}
type structb28 struct {
	c [28]byte
}
type structb29 struct {
	c [29]byte
}
type structb44 struct {
	c [44]byte
}
type structb45 struct {
	c [45]byte
}
type structb60 struct {
	c [60]byte
}
type structb61 struct {
	c [61]byte
}

func main() {

	fmt.Println("Raw sizeof")
	fmt.Println("sizeof(byte)=", unsafe.Sizeof(*new(byte)))
	fmt.Println("sizeof(char/rune)=", unsafe.Sizeof(*new(rune)))
	fmt.Println("sizeof(i16)=", unsafe.Sizeof(*new(int16)))
	fmt.Println("sizeof(i32)=", unsafe.Sizeof(*new(int32)))
	fmt.Println("sizeof(i64)=", unsafe.Sizeof(*new(int64)))
	fmt.Println("sizeof(int)=", unsafe.Sizeof(*new(int)))
	fmt.Println("sizeof(uintptr)=", unsafe.Sizeof(*new(uintptr)))
	fmt.Println("sizeof(float32)=", unsafe.Sizeof(*new(float32)))
	fmt.Println("sizeof(float64)=", unsafe.Sizeof(*new(float64)))

	fmt.Printf("\n\nMeasure allocation\n")
	measure[int]("int")
	measure[structb]("structb")
	measure[structib]("structib")
	measure[structip]("structip")
	measure[structiib]("structiib")
	measure[structbib]("structbib")
	measure[structiii]("structiii")
	measure[structi16b]("structi16b")
	measure[structi32b]("structi32b")
	measure[structb12]("structb12")
	measure[structb13]("structb13")
	measure[structb28]("structb28")
	measure[structb29]("structb29")
	measure[structb44]("structb44")
	measure[structb45]("structb45")
	measure[structb60]("structb60")
	measure[structb61]("structb61")
}
