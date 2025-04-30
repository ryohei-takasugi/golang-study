package main

import (
	"fmt"
	"image"
	"time"
)

func add(x int, y int) int {
	return x + y
}
func add2(x, y int) int {
	return x + y
}
func claculate() {
	fmt.Println(add(42, 13))
	fmt.Println(add2(42, 13))
}

func swap(x, y string) (string, string) {
	return y, x
}
func multiReturn() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func namedReturn() {
	fmt.Println(split(17))
}

func channel() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("sending to c1")
		c1 <- "one"
	}()
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("sending to c2")
		c2 <- "two"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received msg1:", msg1)
		case msg2 := <-c2:
			fmt.Println("received msg2:", msg2)
		}
	}
}

func rangeSample1() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
func rangeSample2() {
	pow2 := make([]int, 10)
	for i := range pow2 {
		pow2[i] = 1 << uint(i)
	}
	fmt.Println(pow2)
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}
}
func deferSample() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

type Vertex struct {
	Lat, Long float64
}

func mapLiteral() {
	var m = map[string]Vertex{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)
}

func mutatingMaps() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func makeCounter() (func() int, func() int) {
	count := 0
	increment := func() int {
		count++
		return count
	}
	decrement := func() int {
		count--
		return count
	}
	return increment, decrement
}

func closures() {
	increment, decrement := makeCounter()
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
	fmt.Println(decrement()) // 1
}

func imageSample() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

func main() {
	deferSample()
	fmt.Println("===================================")
	claculate()
	fmt.Println("===================================")
	multiReturn()
	fmt.Println("===================================")
	namedReturn()
	fmt.Println("===================================")
	channel()
	fmt.Println("===================================")
	rangeSample1()
	fmt.Println("===================================")
	rangeSample2()
	fmt.Println("===================================")
	mapLiteral()
	fmt.Println("===================================")
	mutatingMaps()
	fmt.Println("===================================")
	closures()
	fmt.Println("===================================")
	imageSample()
}
