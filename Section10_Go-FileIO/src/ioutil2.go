package main

import (
	"bufio"
	"fmt"
	"os"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	/*
	      type Reader interface {
	        Read(p []byte) (n int, err error)
	    }
	     type Writer interface {
	        Write(p []byte) (n int, err error)
	   }
	*/

	file, err := os.OpenFile("sample.txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777))

	wt := bufio.NewWriter(file)
	wt.WriteString("Hello Golang1!\nFile Write Test1!\n")
	wt.Write([]byte("Hello Golang2!\nFile Write Test2!\n"))

	errCheck(err)

	fmt.Println(wt.Buffered())
	fmt.Println(wt.Available())
	fmt.Println(wt.Size())

	wt.Flush()

	rt := bufio.NewReader(file) //Reader 반환
	fi, err := file.Stat()
	errCheck(err)

	b := make([]byte, fi.Size())

	fmt.Println(fi)
	fmt.Println(fi.Name())
	fmt.Println(fi.Size())
	fmt.Println(fi.ModTime())

	file.Seek(0, os.SEEK_SET)
	data, _ := rt.Read(b)

	fmt.Println(rt.Size())
	fmt.Println(data)

	fmt.Println(string(b))

	defer file.Close()
}
