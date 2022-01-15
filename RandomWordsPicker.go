package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	file, err := os.Open("words.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	str := buf.String()

	rand.NewSource(time.Now().UnixNano())

	var words []string
	var print string

	words = strings.Split(str, "\n")

	fmt.Println("How many words would you like to generate?")
	var input, i, random int
	fmt.Scanf("%d", &input)
	for i = 0; i < input; i++ {
		random = rand.Intn(len(words) - 0)
		print = words[random]
		if i == 0 {
			print = strings.Title(strings.ToLower(print))
		} else {
			print = strings.ToLower(print)
		}
		fmt.Printf("%s ", print)
	}
	fmt.Println()
	os.Exit(0)
}
