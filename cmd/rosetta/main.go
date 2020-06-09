package main

import (
	"fmt"
	"os"

	flag "github.com/ogier/pflag"

	"github.com/human-caching/rosetta/cmdparser"

	"github.com/human-caching/rosetta/filehdlr"

	"bufio"
)

var context = flag.StringP("context", "c", "", "Context for aliases")

var doInstall = false

func main() {
	if doInstall {
		fmt.Println("Configuration file not found, Initial installation is needed to use Rosetta")
		fmt.Print("Would you like to install Rosetta (Y/N)? ")
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()

		if err != nil {
			fmt.Println("Something went wrong reading your input")
			os.Exit(1)
		}

		install(char)
	} else {
		// Parse commands sent if install is uneeded
		flag.Parse()
	}

}

func init() {
	_, newInstall := cmdparser.ContextCheck()

	doInstall = !newInstall
}

func install(install rune) {
	switch install {
	case 'Y':
		fmt.Println("Installing.....")
		filehdlr.CreateConfigurationFile()
		fmt.Println("Done!")
		break
	case 'y':
		fmt.Println("Installing.....")
		filehdlr.CreateConfigurationFile()
		fmt.Println("Done!")
		break
	default:
		fmt.Println("Ok, no install")
	}

}
