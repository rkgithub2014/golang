package main

import (
	"fmt"
)

type Greeter interface {
	Greet(name string) string
}

func FormatGreeter(format string) Greeter {
	return formatGreeter{Format: format}
}

type formatGreeter struct {
	Format string
}

func main() {
	fmt.Println("main hello")
	var g Greeter
	fmt.Println("Message:", g.Greet("%s hello"))

}

func (g formatGreeter) Greet(name string) string {
	if g.Format == "" {
		g.Format = "Hi %s"
	}
	return fmt.Sprintf(g.Format, name)
}
