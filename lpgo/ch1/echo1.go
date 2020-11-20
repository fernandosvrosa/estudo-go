package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	inicio := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(os.Args[0])
	fmt.Println(s)
	fmt.Println(time.Since(inicio))
}
