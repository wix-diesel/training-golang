package main
import "fmt"

func main(){
	const lightSpeed = 299792 
	var distance = 56000000
	fmt.Println(distance / lightSpeed, "秒")

	distance = 401000000
	fmt.Println(distance/lightSpeed, "秒")
}