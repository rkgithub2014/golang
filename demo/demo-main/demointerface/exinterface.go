package main

import (
	"errors"
	"flag"
	"fmt"
)

func main() {
	dbCmd := command{"db", "create", make(map[string]string)}
	args := []string{"db", "create", "-dbtype", "moodysdb", "-validate", "pkdir=/usr/18.22"}
	execute(&dbCmd, args[2:])
}

type ICommand interface {
	execute()
}

/*
type iClCmd interface {
	InitFlags(args []string)
	ValidateFlags() error
}
*/

var dbtype string
var validate bool
var dbCreateFlagSet = flag.NewFlagSet("dbCreateFlagSet", flag.ContinueOnError)

func execute(clcmd *command, args []string) {
	clcmd.InitFlags(args)
	err := clcmd.ValidateFlags()
	if err == nil {
		fmt.Println("All flags set, call SDK function")
	}
}

type command struct {
	cmd    string
	subcmd string
	args   map[string]string
}

func (dbcmd *command) ValidateFlags() error {
	if dbtype == "" {
		msg := dbCreateFlagSet.Lookup("dbtype").Usage
		fmt.Println(msg)
		err := errors.New(msg)
		return err
	}
	return nil
}

func (dbcmd *command) InitFlags(args []string) {
	//dbCreateFlagSet := flag.NewFlagSet("dbCreateFlagSet", flag.ContinueOnError)
	//dbtype := dbCreateFlagSet.String("dbtype", "", "dbtype: moodysdb|tenantdb|olapdb (Required)")
	dbCreateFlagSet.StringVar(&dbtype, "dbtype", "", "-dbtype: moodysdb|tenantdb|olapdb (Required)")
	dbCreateFlagSet.BoolVar(&validate, "validate", false, "validate without commit")
	dbCreateFlagSet.Parse(args)
	//fmt.Println("args", args)
	//fmt.Println("Parsed dbtype", dbtype)
	//fmt.Println("Parsed validate", validate)
}
