package utils

import "fmt"

// Log - consistenet logging
func Log(msg interface{}) {
	fmt.Printf("\n%v\n", msg)
}
