package main

import (
	"github.com/fogleman/gg"
	"path/filepath"
	"fmt"
	"math"
)

func mySin(t float64) float64{
	return 50*math.Sin(t/10) + 350
}
func DrawBasicShape(){
	dc := gg.NewContext(1000, 700)
	dc.Push()
	dc.SetLineWidth(2)
	dc.SetRGB(10, 80, 24)
	dc.DrawLine(20,350,980,350)
	dc.DrawLine(500,200,500,500)
	dc.Stroke()
	dc.Pop()
	dc.MoveTo(40,350)
	amount :=1000.0
	step:=920.0/amount
	dc.SetLineWidth(1.5)
	dc.SetRGB(100, 10, 20)
	for i := 0.0;i<=amount;i+=0.1 {
		x:=498+(i-amount/2)*step
		y:=mySin(i)
		dc.LineTo(x,y)
		//dc.MoveTo(x,y)
	}
	dc.Stroke()
	imgDir,_ := filepath.Abs("./iris-web/assets/img")
	filePath:= filepath.Join(imgDir,"circle.png")
	fmt.Println(filePath)
	dc.SavePNG(filePath)
}
func main() {
	DrawBasicShape()
}