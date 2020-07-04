package dbcreate

import (
	"demo/globals"
	"flag"
	"fmt"
	"strconv"
)

// define flagset, flag
var dbCreateFlagSet *flag.FlagSet
var dbtype string
var validate bool

// define command type instance
type dbCmdType globals.CmdType

// implement globals.CmdHandler{} Execute() method
func (dbcmd dbCmdType) Execute() error {
	fmt.Println("execute():", dbcmd.Cmd)
	return nil
}

// implement InitFlags() method
func (dbcmd dbCmdType) InitFlags(args []string) {
	fmt.Println("Initflag():\n Input Args", args)
	dbCreateFlagSet = flag.NewFlagSet(globals.DbCreateFlagSet, flag.ContinueOnError)
	dbCreateFlagSet.StringVar(&dbtype, globals.DbtypeFlag, "", globals.DbtypeFlagUsage)
	dbCreateFlagSet.BoolVar(&validate, globals.ValidateFlag, false, globals.ValidateFlagUsage)
	dbCreateFlagSet.Parse(args)
}

// implement Validate() method
func (dbcmd dbCmdType) Validate() error {
	fmt.Println("Validate():", dbcmd.Cmd)
	// Required flag
	if dbtype == globals.EmptyString {
		return globals.ErrdbtypeFlagMissing
	}
	// TODO - validate correct value
	// Add to arguments
	dbcmd.CmdArgs[globals.DbtypeFlag] = dbtype
	dbcmd.CmdArgs[globals.ValidateFlag] = strconv.FormatBool(validate)
	fmt.Println("Parsed Flags", dbcmd.CmdArgs)
	return nil
}

//CallHandle -- use to delegate to command
func CallHandle(args []string, cmdinst globals.CmdType) error {
	execCmd := todbCmdType(cmdinst)
	err := handle(args, execCmd)
	// Call Sdk function
	if err == nil {
		err = execCmd.Execute()
	}
	return err
}

func handle(args []string, cmdhandler globals.CmdHandler) error {
	//Init flags
	cmdhandler.InitFlags(args)
	//Validate flags
	return cmdhandler.Validate()
}

func todbCmdType(cmdinst globals.CmdType) dbCmdType {
	return dbCmdType{cmdinst.Cmd, cmdinst.SubCommand, cmdinst.CmdArgs}
}
