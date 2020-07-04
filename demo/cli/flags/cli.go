package main

import (
	"demo/cmdhandler"
	"demo/globals"
	"errors"
	"flag"
	"fmt"
	"strings"
)

func main() {

	// init flags
	cmdhandler.InitFlags()

	if flag.NArg() == 0 || flag.NArg() < 2 {
		fmt.Println("Display Usage")
		return
	}

	// Process input
	cmdinst, err := ParseArgs()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Display Usage")
		return
	}
	switch cmdinst.Cmd {
	case "db":

	}
}

// ParseArgs input arguments
func ParseArgs() (globals.CmdType, error) {
	args := flag.Args()
	var cmdinst = globals.CmdType{}
	// parse cmd, subcommand
	cmdinst.Cmd = args[0]
	cmdinst.SubCommand = args[1]
	if cmdinst.SubCommand == "" {
		err := errors.New("subcommand missing")
		return cmdinst, err
	}
	// parse args map[string]string param=value
	var inputArgs = make(map[string]string)
	for i := 2; i < len(args); i++ {
		if strings.Index(args[i], "=") != -1 {
			arginst := strings.Split(args[i], "=")
			inputArgs[arginst[0]] = arginst[1]
			cmdinst.CmdArgs = inputArgs
		}
	}
	if len(cmdinst.CmdArgs) == 0 {
		err := errors.New("arguments not found")
		return cmdinst, err
	}
	fmt.Println("{\nCommand:", cmdinst.Cmd)
	fmt.Println("SubCommand:", cmdinst.SubCommand)
	fmt.Println("Parsed args:", cmdinst.CmdArgs)
	fmt.Println("}")
	return cmdinst, nil
}
