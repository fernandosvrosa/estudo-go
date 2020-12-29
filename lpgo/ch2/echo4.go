// Echo4 exibe seus argumentos de linha de comnado
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separetor")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

/*\
casos de testes
go build echo4.go
./echo4  a bc def
./echo4 -s / a bc def
./echo4 -n a bc def
./echo4 -help
 */
