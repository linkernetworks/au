package client

import "fmt"

//printResult - display json to table in console
func printResult(ret []byte) {
	// Using standard first, will use kubernetes printer
	fmt.Println(string(ret))
}
