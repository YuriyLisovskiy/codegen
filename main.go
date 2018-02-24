package main

import (
	"os"
	"./execute"
)

func main() {
	execute.ExecuteFromCommandLine(os.Args[1:])
}
