package main

import (
	"fmt"
	"strconv"

	"github.com/human-caching/rosetta/cmdparser"
)

func main() {
	path, newInstall := cmdparser.ContextCheck()

	fmt.Println("Current Config Directory: " + path)
	fmt.Println("Is this a new install: " + strconv.FormatBool(!newInstall))
}
