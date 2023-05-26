package main

import (
	"fmt"

	"github.com/fahadsiddiqui/gochunker/chunker"
)

func main() {
	chnks, _ := chunker.Chunker("hey, what's up what are you doing these days", 12)
	for _, ch := range chnks {
		fmt.Println(ch)
	}
}
