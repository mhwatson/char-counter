package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	filePtr := flag.String("file", "proverbs.txt", "the path to the file.")
	flag.Parse()

	file := *filePtr

	fileByte, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	fileInStr := string(fileByte)
	lines := strings.Split(fileInStr, "\n")
	for _, l := range lines {
		fmt.Printf("%s\n", l)
		for k, v := range charCount(l) {
			fmt.Printf("'%c'=%d, ", k, v)
		}
		fmt.Print("\n\n")
	}
}

func charCount(line string) map[rune]int {
	m := make(map[rune]int, 0)
	for _, c := range line {
		m[c] = m[c] + 1
	}
	return m
}
