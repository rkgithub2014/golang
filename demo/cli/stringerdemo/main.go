package main

import "fmt"

type line struct {
	start  int
	end    int
	dotted bool
}

func (l line) String() string {
	return fmt.Sprintf("Line start=%d,end=%d,dotted=%t", l.start, l.end, l.dotted)
}

func main() {
	var dottedLine = line{start: 1, end: 10, dotted: true}
	var newline line
	fmt.Println("dottedline", dottedLine)
	fmt.Println("newline", newline)

}
