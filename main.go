package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		log.Panicf("Create() failed: %v", err.Error())
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for i := 0; i != 10000000; i++ {
		_, err := w.WriteString(fmt.Sprintf("Line %v (zero-base index: %v)\r\n", i+1, i))
		if err != nil {
			log.Panicf("WriteString() failed: %v", err.Error())
		}
	}
	w.Flush()

}
