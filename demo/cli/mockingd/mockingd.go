package mockingd

import (
	"fmt"
	"io"
	"time"
)

//Countdown prints a countdown from 3 to out
func Countdown(out io.Writer) {
	fmt.Fprintf(out, "3")
}

const finalWord = "Go!"
const countdownStart = 3

func CountDownConst(out io.Writer) {

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, finalWord)
}

//Sleeper allows to put a delay
type Sleeper interface {
	Sleep()
}

//DefaultSleeper is default implementation with fixed delay
type DefaultSleeper struct{}

// Sleep will pause execution for defined duration
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// CountDownDelay function uses default sleeper to pause execution
func CountDownDelay(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}
