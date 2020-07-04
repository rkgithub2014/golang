package cmdhandler

import (
	"demo/globals"
	"flag"
	"fmt"
)

//InitFlags initialize all flags
func InitFlags() {
	// initialize
	var dbType string
	flag.StringVar(&dbType, globals.DbtypeFlag, "", "Usage: moodysdb|tenantdb|olapdb")
	var ftType string
	flag.StringVar(&ftType, globals.FttypeFlag, "", "Usage: all|single|standard")
	flag.Parse()
}

//Execute handles DB commands
func Execute(cmdinst globals.CmdType) {

	switch cmdinst.Cmd {
	case "create":
		CheckFlags()
	}
	// db create
	// db upgrade
	// db install
	// db copy
	// db delete
}

func CheckFlags() {
	fmt.Print("CheckFlags")
}
