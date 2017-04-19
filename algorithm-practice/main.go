package main

import (
	"fmt"
	//"unicode/utf8"
	//"math"
	//"math/cmplx"
	//"math/rand"
	"runtime"
	//"strconv"
	//"strings"
	"time"
	alg "./algorithm"
)
func main(){
	t := time.Now()
	fmt.Printf("%4d-%02d-%02d %02d:%02d:%02d.%d\n ", t.Year(),
		t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond())
	fmt.Printf("%s\t%s\n", runtime.GOOS, runtime.Version())
	/*arr := []int{6, 8, 3, 2, 7, 5, 4, 9}
	for i := 0; i < 1000000; i++ {
		arr = append(arr, int(rand.Int31n(1000)))
	}*/
	start := time.Now()
	for i:=1;i<30;i++{
		sli:=alg.YangTriangle(i)
		fmt.Println(sli)
	}
	end := time.Now()
	fmt.Println("耗时：", end.Sub(start))

	rangeParams("sting",14)
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024)
}

func rangeParams(args ...interface{}){
	fmt.Printf("%T %T\n",args[0],args[1])
	fmt.Printf("%v\n",args)
}
