package main

import (
	"fmt"
	"time"
)

func LongProcessCalc(x, y int) int {
	time.Sleep(5 * time.Second)
	return x + y
}

type TimeInterface interface {
	Sleep(d time.Duration)
}

func main() {
	result := LongProcessCalc(2, 3)
	fmt.Println(result)
}
