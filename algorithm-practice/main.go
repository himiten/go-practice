package main

import (
	"fmt"
	//"unicode/utf8"
	//"math"
	//"math/cmplx"
	"math/rand"
	"runtime"
	//"strconv"
	//"strings"
	"time"
	algo "./algorithm"
)
func main(){
	t := time.Now()
	fmt.Printf("%4d-%02d-%02d %02d:%02d:%02d.%d\n ", t.Year(),
		t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond())
	fmt.Printf("%s\t%s\n", runtime.GOOS, runtime.Version())
	arr := []int{6, 8, 3, 2, 7, 5, 4, 9}
	for i := 0; i < 1000000; i++ {
		arr = append(arr, int(rand.Int31n(1000)))
	}
	start := time.Now()
	algo.QuickSort(arr)
	end := time.Now()
	fmt.Println("耗时：", end.Sub(start))
}