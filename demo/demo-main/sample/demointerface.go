package main

import (
	"fmt"
)

func main() {
	fmt.Println("Test")
	s1 := square{10, 10}
	calc(s1)

	// command instance
	dbcmd := clcmd{"db", "upgrade"}
	// execute command
	handle(dbcmd)
}

type geometry interface {
	area() float64
}

type square struct {
	length float64
	width  float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func calc(g geometry) {
	fmt.Println("Calculated Area:", g.area())
}

type clcmdhandler interface {
	initflag()
	validate() error
	execute()
}

type clcmd struct {
	cmd    string
	subcmd string
}

func handle(cmdhandler clcmdhandler) {
	cmdhandler.initflag()
	cmdhandler.validate()
	cmdhandler.execute()
}

func (dbcmd clcmd) execute() {
	fmt.Println("execute method", dbcmd.cmd, dbcmd.subcmd)
}

func (dbcmd clcmd) initflag() {
	fmt.Println("initflag method", dbcmd.cmd)
}

func (dbcmd clcmd) validate() error {
	fmt.Println("validate method", dbcmd.cmd)
	return nil
}
