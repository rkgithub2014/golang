package globals

import (
	"fmt"
	"os"
	"strings"
)

//CliParseArgs - Parses Command, SubCommand and Arguments
func CliParseArgs(args []string) (CmdType, error) {
	var cmdinst = CmdType{}
	// parse cmd, subcommand
	cmdinst.Cmd = args[1]
	cmdinst.SubCommand = args[2]
	if cmdinst.SubCommand == "" {
		err := ErrSubCommandMissing
		return cmdinst, err
	}
	// parse args map[string]string param=value
	var inputArgs = make(map[string]string)
	for i := 3; i < len(args); i++ {
		if strings.Index(args[i], "=") != -1 {
			arginst := strings.Split(args[i], "=")
			inputArgs[arginst[0]] = arginst[1]
			cmdinst.CmdArgs = inputArgs
		}
	}
	if len(cmdinst.CmdArgs) == 0 {
		return cmdinst, ErrArgsNotFound
	}
	return cmdinst, nil
}

//CliCheckError - Check Error in Argument Parsing
func CliCheckError(err error) {
	if err != nil && (err == ErrArgsNotFound || err == ErrSubCommandMissing) {
		fmt.Println(err)
		usage()
		os.Exit(1)
	}
}

//CliCheckArgs - Check Arguments
func CliCheckArgs(args []string) error {
	if len(args) == 0 || len(args) < 2 {
		return ErrArgsNotFound
	}
	return nil
}

func usage() {
	fmt.Println("Usage: cldeploy <cmd> <subcommand> [options] <parameters>")
}
