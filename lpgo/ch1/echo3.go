package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	inicio := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(time.Since(inicio))
}
