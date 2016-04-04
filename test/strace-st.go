package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	// "time"
)

func readFile2(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	fmt.Println(string(fd))
	return string(fd)
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func main() {

	flag.Parse()
	inputFile := flag.Arg(0)
	outputFile := flag.Arg(1)
	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(nil)
		os.Exit(1)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	f1, _ := os.OpenFile(outputFile, os.O_APPEND, 0666)
	if checkFileIsExist(outputFile) {
		//f, err1 := os.OpenFile(outputFile, os.O_APPEND, 0666)
		os.Remove(outputFile)
		f1, _ = os.Create(outputFile)
		// fmt.Println("文件存在")
	} else {
		//f, err1 := os.Create(outputFile)
		f1, _ = os.Create(outputFile)
		fmt.Println("OutputFile not found.")
	}

	for scanner.Scan() {
		// fmt.Println(strings.TrimSpace(scanner.Text()))
		buffString := strings.TrimSpace(scanner.Text())
		// fmt.Println(buffString)
		// i := strings.IndexAny(buffString, "(")
		// fmt.Println(i)
		buffA := strings.Split(buffString, "(")
		buffB := strings.Split(buffA[0], " ")
		//fmt.Println(buffB[1])
		//check(err1)
		//n, err1 := io.WriteString(f, buffB[1])
		io.WriteString(f1, buffB[1])
		io.WriteString(f1, "\n")
		//check(err1)
	}

	// fmt.Printf("%d bytes writed", n)
}
