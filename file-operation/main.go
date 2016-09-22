package main

import (
	//"bufio"
	"fmt"
	//"io"
	"io/ioutil"
	"os"
)

func main() {
	/*var name string
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&name)
	fmt.Println("hello",name)
	var t1 float32
	var t2 int
	var t3 string
	fmt.Sscanf("12.5,40,magic","%f,%d,%s",&t1,&t2,&t3)
	fmt.Println("From the string we read: ",t1, t2, t3)*/
	/*var inputReader *bufio.Reader
	var input string
	var err error

	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}*/
	inputFile, inputErr := os.Open("public/font.scss")
	if inputErr != nil {
		fmt.Printf("文件读取失败:\n%v\n",inputErr)
		return
	}
	defer inputFile.Close()
	/*inputReader := bufio.NewReader(inputFile)
	lineNum:=1*/
	/*for {
		inputStr, readErr := inputReader.ReadString('\n')
		if readErr == io.EOF{
			break
		}
		fmt.Printf("%d:%s\n",lineNum,inputStr)
		lineNum++
	}*/
	buf,err:=ioutil.ReadFile("public/font.scss")
	if err!=nil{
		fmt.Fprintf(os.Stderr,"File Error:%s\n",err)
	}
	//fmt.Printf("%v\n", buf)
	//fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile("temp/test.txt", buf, 0644) // oct, not hex
	if err != nil {
		panic(err. Error())
	}

}
