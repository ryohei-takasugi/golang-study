package main

import (
	"fmt"
	"math"
	"unsafe"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	fmt.Println("X  :", v.X)
	fmt.Println("Y  :", v.Y)
	fmt.Println("X*X:", v.X*v.X)
	fmt.Println("Y*Y:", v.Y*v.Y)
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func method1() {
	v := Vertex{3, 4}
	fmt.Println("abs:", v.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func method2() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

type Vertex2 struct {
	X, Y float64
}

func (v Vertex2) Abs() float64 {
	fmt.Println("X  :", v.X)
	fmt.Println("Y  :", v.Y)
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex2) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func method3() {
	v := Vertex2{3, 4}
	fmt.Println("abs:", v.Abs())
	v.Scale(10)
	fmt.Println("abs:", v.Abs())
}

func Float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

func Float64frombits(b uint64) float64 {
	return *(*float64)(unsafe.Pointer(&b))
}

func Abs(x float64) float64 {
	return Float64frombits(Float64bits(x) &^ (1 << 63))
}

func floatSample() {
	x := -3.5
	bits := Float64bits(x)
	fmt.Printf("bits = %064b\n", bits)

	disiableBits := uint64(1) << 63
	fmt.Printf("bits = %064b\n", disiableBits)

	formBits := bits &^ disiableBits
	fmt.Printf("bits = %064b\n", formBits)

	reconstructed := Float64frombits(formBits)
	fmt.Printf("reconstructed = %f\n", reconstructed)

	absValue := Abs(x)
	fmt.Printf("absValue = %f\n", absValue)
}

func pointerSample() {
	var a int = 10
	var b *int
	b = &a
	fmt.Printf("a = %d\n", a)
	fmt.Printf("b = %d\n", *b)
}

func main() {
	method1()
	fmt.Println("==========================")
	method2()
	fmt.Println("==========================")
	method3()
	fmt.Println("===========================")
	floatSample()
	fmt.Println("===========================")
	pointerSample()
	fmt.Println("===========================")
}
