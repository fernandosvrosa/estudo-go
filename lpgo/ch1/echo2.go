package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	inicio := time.Now()
	for indice, arg := range os.Args[1:] {
		fmt.Println(indice, arg)
	}
	fmt.Println(time.Since(inicio))

}
