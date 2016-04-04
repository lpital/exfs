package main

import (
	"bufio"
	"flag"
	"fmt"
	// "io"
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

func main() {

	flag.Parse()
	file := flag.Arg(0)
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(nil)
		os.Exit(1)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(strings.TrimSpace(scanner.Text()))
	}
}
