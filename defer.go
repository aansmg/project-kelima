// package main

// import "fmt"

// func main() {
// 	defer fmt.Println("defer function starts to execute")
// 	fmt.Println("Hey everyone")
// 	fmt.Println("Welcome Back to Go learning centre")
// }

// func main() {
// 	callDeferFunc(false)
// 	fmt.Println("Hey Everyone")
// }

// func callDeferFunc(condition bool) {
// 	if condition == false {
// 		fmt.Println("Dont call deferfunc")

// 		return
// 	}
// 	defer deferFunc()
// }
// func deferFunc() {
// 	fmt.Println("Defer func stats to execute")
// }

package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Invoke with defer")
	fmt.Println("Before exit")
	os.Exit(1)
}
