package main

import (
	"github.com/anriclee/GoPatterns/cmdpattern"
	"os"
)

func main() {
	cmdpattern.Execute(os.Args[1:])
}
