package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var randomNumber = rand.Intn(100) + 1
	fmt.Println("1~100の愛でで文字を当ててください")

	for {
		var answer = rand.Intn(100) + 1
		fmt.Println("コンピューターは", answer, "を選びました。")

		if answer == randomNumber {
			fmt.Println("正解！")
			break
		} else if answer < randomNumber {
			fmt.Println("小さすぎます！")
		} else {
			fmt.Println("大きすぎます！")
		}
	}
}
