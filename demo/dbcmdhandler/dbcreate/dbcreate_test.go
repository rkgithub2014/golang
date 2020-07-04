package dbcreate

import (
	"demo/globals"
	"testing"
)

func TestInvalidcmd(t *testing.T) {
	// test case arguments
	cmdArgs := []string{"cldeploy"}
	tt := struct {
		name     string
		args     []string
		expected error
	}{"InvalidCmdTest", cmdArgs, globals.ErrArgsNotFound}
	//t.Skip("skipping", tt.name)
	t.Run(tt.name, func(t *testing.T) {
		ok := globals.CliCheckArgs(tt.args)
		if ok != tt.expected {
			t.Errorf(ok.Error())
		}
	})
}
func TestHandle(t *testing.T) {
	// test case arguments
	type testcaseArg struct {
		args       []string
		cmdhandler globals.CmdHandler
	}
	cmdArgs := []string{"cldeploy", "db", "create", "-dbtype", "moodysdb", "pkgdir=/usr/local/18.22", "config=/usr/local/cldeploy.json"}
	// cli cmd validation
	ok := globals.CliCheckArgs(cmdArgs)
	if ok != nil {
		t.Errorf(ok.Error())
	}
	// cli parse arguments
	cmdinst, ok := globals.CliParseArgs(cmdArgs)
	if ok != nil {
		t.Errorf(ok.Error())
	}
	// call dbcreate.handle
	argsinst := testcaseArg{cmdArgs[3:], todbCmdType(cmdinst)}
	tests := []struct {
		name string
		args testcaseArg
	}{
		// TODO: Add test cases.
		{"dbcreateTest", argsinst},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok := handle(tt.args.args, tt.args.cmdhandler)
			if ok != nil {
				t.Errorf(ok.Error())
			}
		})
	}
}
