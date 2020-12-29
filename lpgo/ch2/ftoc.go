// Ftoc exibe duas conversões de Fahrenheit para Celsius
package main

import "fmt"

func main() {
	const freezzingF , boilingF = 32.0, 212.0
	fmt.Printf("%g F = %g C\n", freezzingF, fToC(freezzingF))
	fmt.Printf("%g F = %g C\n", boilingF, fToC(boilingF))
}

func fToC(f float64)float64 {
	return (f - 32) * 5 /9
}
