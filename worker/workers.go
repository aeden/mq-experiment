package main

import (
	"runtime"
)

func main() {
	runLogger()
	runResponder()
	runtime.Goexit()
}
