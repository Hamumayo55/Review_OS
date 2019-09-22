package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var BufferSize, WriteSize int = 1024, 102400
var buffering bool = false

func help() {
	fmt.Printf("-b -> buffered Read\n")
	fmt.Printf("-u -> unbuffered Read\n")
	fmt.Printf("--size 1024 -> Read file size\n")
}

func getopts(args []string) {
	if len(args) == 0 {
		help()
	}
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "b":
			buffering = true
			if new_bs, err := strconv.Atoi(args[i+1]); err == nil {
				BufferSize = new_bs
				i = i + 1
				fmt.Printf("new buffer size is %d\n", BufferSize)
			}
		case "u":
			buffering = false
			fmt.Printf("unbuffer\n")
		case "s":
			if new_ws, err := strconv.Atoi(args[i+1]); err == nil {
				WriteSize = new_ws
				i = i + 1
				fmt.Printf("new Read size is %d\n", WriteSize)
			}
		default:
			help()
		}
	}
}

func FileRead() {
	FILE_PATH := `PATH/To/File`
	file, err := os.Create(FILE_PATH)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	output := "0"

	if buffering {
		buffer := bufio.NewWriterSize(file, BufferSize)

		for i := 0; i < WriteSize; i++ {
			buffer.Write(([]byte)(output))
		}
		buffer.Flush()
		fmt.Printf("buffer file write finished\n")
	} else {
		for i := 0; i < WriteSize; i++ {
			file.Write(([]byte)(output))
		}
		fmt.Printf("unbuffer file write finished\n")
	}
}

func main() {
	flag.Parse()
	args := flag.Args()

	getopts(args)

	start := time.Now()
	FileRead()
	end := time.Now()

	fmt.Printf("%f\n", (end.Sub(start)).Seconds())
}
