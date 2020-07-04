package main

import (
	"demo/dbcmdhandler/dbcreate"
	"demo/globals"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	// Validate Args
	ok := globals.CliCheckArgs(args)
	globals.CliCheckError(ok)
	// Parse Args
	cmdinst, ok := globals.CliParseArgs(args)
	globals.CliCheckError(ok)
	// call handler
	switch cmdinst.Cmd {
	case "db":
		ok = dbcreate.CallHandle(args[3:], cmdinst)
	}
	if ok != nil {
		fmt.Println(ok)
		os.Exit(1)
	}
	printCmd(&cmdinst)
}

func printCmd(cmdinst *globals.CmdType) {
	fmt.Println("{\nCommand:", cmdinst.Cmd)
	fmt.Println("SubCommand:", cmdinst.SubCommand)
	fmt.Println("Parsed args:", cmdinst.CmdArgs)
	fmt.Println("}")

}
