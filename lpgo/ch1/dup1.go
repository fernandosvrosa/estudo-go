// Dup1 exibe o texto de toda linha que aparece mais de
// uma vez na entrada-padrao, precedida por sua contagem.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	conts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		if input.Text() == "" {
			break
		}
		conts[input.Text()]++
	}
	// NOTA: ignorando erros em potencial de input.Err()
	for line, n := range conts{
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}



