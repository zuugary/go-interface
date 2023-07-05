package main

import (
	"fmt"
	"time"
)

type Clock interface {
	Sleep(d time.Duration)
}

func LongProcessCalc(x, y int, t Clock) int {
	t.Sleep(5 * time.Second)
	return x + y
}

type realClock struct{}

func (r *realClock) Sleep(d time.Duration) {
	time.Sleep(d)
}

func main() {
	result := LongProcessCalc(2, 3, &realClock{})
	fmt.Println(result)
}
