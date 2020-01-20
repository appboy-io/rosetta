package main

import (
	flag "github.com/ogier/pflag"

	"github.com/human-caching/rosetta/cmdparser"
)

var context = flag.StringP("context", "c", "", "Context for aliases")

var doInstall = false

func main() {
	if doInstall {
		println("Initial installation is needed to use Rosetta")
	} else {
		// Parse commands sent if install is uneeded
		flag.Parse()
	}

}

func init() {
	_, newInstall := cmdparser.ContextCheck()

	doInstall = !newInstall
}
