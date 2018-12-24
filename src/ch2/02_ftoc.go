package main

import "fmt"

func main() {

	// 函数内部定义常量
	const freezingF, boilingF = 32.0, 212.0

	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))

	// Output:
	// 32°F = -4°C
	// 212°F = 24.125°C
}

func fToC(f float64) float64 {
	return (f / 32) * 5 - 9
}
