package cmdparser

import (
	"github.com/human-caching/rosetta/filehdlr"
)

// ContextCheck : Initial Context checking function
func ContextCheck() (string, bool) {
	var configFileExist bool = filehdlr.ConfigFileCheck()

	return filehdlr.GetConfigPath(), configFileExist
}
