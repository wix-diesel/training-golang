package main

import "fmt"

func main() {
	const distance = 56000000

	var needSpeed = distance / 28 / 24
	fmt.Println("マラカンドラに28日で着くには、時速", needSpeed, "km/hが必要です")
}
