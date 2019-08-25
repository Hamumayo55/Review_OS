package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func make_symbol() string {
	var ch_slice = []string{}

	for n := 0; n < 10; n++ {
		asciiValue_number := 48 + n
		ch_slice = append(ch_slice, string(asciiValue_number))
	}

	for j := 0; j < 26; j++ {
		asciiValue_large_symbol := 65 + j
		ch_slice = append(ch_slice, string(asciiValue_large_symbol))
	}

	for i := 0; i < 26; i++ {
		asciiValue_small_symbol := 97 + i
		ch_slice = append(ch_slice, string(asciiValue_small_symbol))
	}
	return choice(ch_slice)
}

func choice(s []string) string {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(s))
	return s[i]
}

func main() {
	PassWord := []string{}

	flag.Parse()
	args := flag.Args()
	len_pw, _ := strconv.Atoi(args[0])

	for m := 0; m < len_pw; m++ {
		PassWord = append(PassWord, make_symbol())
	}
	fmt.Println(PassWord)
}
