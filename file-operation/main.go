package main

import (
	//"bufio"
	"fmt"
	//"io"
	"io/ioutil"
	"os"
	//"path/filepath"
	//"bufio"
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
	/*inputFile, inputErr := os.Open("../public/font.scss")
	if inputErr != nil {
		fmt.Printf("文件读取失败:\n%v\n",inputErr)
		return
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	lineNum:=1
	for {
		inputStr, readErr := inputReader.ReadString('\n')
		if readErr == io.EOF{
			break
		}
		fmt.Printf("%d:%s\n",lineNum,inputStr)
		lineNum++
	}*/
	buf,err:=ioutil.ReadFile("../public/font.scss")
	if err!=nil{
		fmt.Fprintf(os.Stderr,"File Error:%s\n",err)
	}
	fmt.Printf("%s\n", string(buf))
	//err = ioutil.WriteFile("../temp/test.txt", buf, 0644) // oct, not hex
	//if err != nil {
	//	panic(err. Error())
	//}
	dirpath:="."
	if len(os.Args) > 1 {
		dirpath=os.Args[1]
	}
	files,err:=ioutil.ReadDir(dirpath)
	if err!=nil {
		fmt.Println(err.Error())
	}else{
		for _,file := range files{
			fmt.Println(file)
			fmt.Printf(" %s  ",file.Name())
		}
	}
	/*err:=filepath.Walk(dirpath, func(path string, f os.FileInfo, err error) error  {
		 if err != nil {
	 		 return err
	 	 }
		 fmt.Printf(" %s  ",path)
		 return  nil
	 })
	 if err !=nil {
		 panic(err.Error())
	 }*/

}
