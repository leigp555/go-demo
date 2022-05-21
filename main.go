package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func testNewReader() {
	r, err := os.OpenFile("a.txt", os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	r2 := bufio.NewReader(r)
	s, _ := r2.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
}
func reset() {
	s1 := strings.NewReader("hello")
	s2 := strings.NewReader("world")
	reader1 := bufio.NewReader(s1)
	readString1, _ := reader1.ReadString('\n')
	fmt.Println(readString1)
	reader1.Reset(s2)
	readString2, _ := reader1.ReadString('\n')
	fmt.Println(readString2)
}
func buff() {
	r, _ := os.OpenFile("a.txt", os.O_RDWR, os.ModePerm)
	newR := bufio.NewReader(r)
	buf := make([]byte, 10)
	for {
		n, err := newR.Read(buf)
		if err == io.EOF {
			break
		} else {
			fmt.Println(string(buf[0:n]))
		}
	}
}
func unReadByte() {
	r, _ := os.OpenFile("a.txt", os.O_RDWR, os.ModePerm)
	r2 := bufio.NewReader(r)
	readByte, _ := r2.ReadByte()
	fmt.Printf("%c\n", readByte)
	readByte, _ = r2.ReadByte()
	fmt.Printf("%c\n", readByte)
	r2.UnreadByte()
	readByte, _ = r2.ReadByte()
	fmt.Printf("%c\n", readByte)
}
func readRune() {
	r := strings.NewReader("你好，我是xxx")
	r2 := bufio.NewReader(r)
	readRune, size, _ := r2.ReadRune()
	fmt.Printf("%c---%v\n", readRune, size)
	readRune, size, _ = r2.ReadRune()
	fmt.Printf("%c---%v\n", readRune, size)
	r2.UnreadRune()
	readRune, size, _ = r2.ReadRune()
	fmt.Printf("%c---%v\n", readRune, size)
}
func readLine() {
	r, _ := os.OpenFile("a.txt", os.O_RDWR, os.ModePerm)
	r2 := bufio.NewReader(r)
	line, isPrefix, _ := r2.ReadLine() //line读取的内容//isPrefix是否读的是前缀
	fmt.Printf("%c\n", line)
	fmt.Println(isPrefix)
}
func readSlice() {
	r, _ := os.OpenFile("a.txt", os.O_RDWR, os.ModePerm)
	r2 := bufio.NewReader(r)
	slice, _ := r2.ReadSlice('\n')
	fmt.Println(string(slice))
}
func readString() {
	r, _ := os.OpenFile("a.txt", os.O_RDWR, os.ModePerm)
	r2 := bufio.NewReader(r)
	content, _ := r2.ReadString('\n')
	fmt.Printf("%v\n", content)
}
func writeTo() {
	r, _ := os.OpenFile("a.txt", os.O_RDWR, os.ModePerm)
	b, _ := os.OpenFile("b.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
	defer r.Close()
	defer b.Close()
	r2 := bufio.NewReader(r)
	buf := bytes.NewBuffer(make([]byte, 10))
	//r2.WriteTo(buf)
	r2.WriteTo(b)
	fmt.Printf("%s", buf)
}

func write() {
	w, _ := os.OpenFile("c.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	defer w.Close()
	w2 := bufio.NewWriter(w)
	w2.Write([]byte("hello world")) //字节形式书写
	w2.WriteString("hello world")   //直接写字符串进去
	w2.Flush()
}
func resetB() {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf2 := bufio.NewWriter(buf)
	buf3 := bytes.NewBuffer(make([]byte, 0))
	buf2.Write([]byte("124455"))
	buf2.Flush()
	buf2.Reset(buf3)
	buf2.WriteString("64656756")
	buf2.Flush()
	fmt.Println(buf)
	fmt.Println(buf3)
}
func available() {
	bf := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(bf)
	fmt.Println(bw.Available())
	fmt.Println(bw.Buffered())
	bw.WriteString("hello")
	fmt.Println(bw.Available())
	fmt.Println(bw.Buffered())
	bw.Flush()
	fmt.Println(bw.Available())
	fmt.Println(bw.Buffered())
}
func rw() {
	read := strings.NewReader("xxx")
	r := bufio.NewReader(read)
	write := bytes.NewBuffer(make([]byte, 0))
	w := bufio.NewWriter(write)
	rw := bufio.NewReadWriter(r, w)
	rw.WriteString("hello")
	rw.Flush()
	s, _ := rw.ReadString('\n')
	fmt.Printf("s:%v\n", s)
	fmt.Printf("buf:%v", write)
}
func readFrom() {
	bf := bytes.NewBuffer(make([]byte, 0))
	s := strings.NewReader("fhggdghdjh")
	nbf := bufio.NewWriter(bf)
	nbf.ReadFrom(s)
	nbf.Flush()
	fmt.Println(bf)
}
func main() {
	s := strings.NewReader("abc def ghi")
	ns := bufio.NewScanner(s)
	ns.Split(bufio.ScanWords)
	for ns.Scan() {
		fmt.Println(ns.Text())
	}

}
