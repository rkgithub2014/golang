package mockingd

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()

	expected := "3"

	if got != expected {
		t.Errorf("got=%s, expected=%s", got, expected)
	}
}

func TestCountDownConst(t *testing.T) {
	buffer := &bytes.Buffer{}
	CountDownConst(buffer)
	got := buffer.String()
	want := `3
2
1
Go!`
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func NewSleeper(calls int) *SpySleeper {
	s := SpySleeper{Calls: calls}
	return &s
}
func TestCountDownDelay(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}
	CountDownDelay(buffer, spySleeper)
	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
	}

	s1 := NewSleeper(1)
	s2 := NewSleeper(2)

	fmt.Println("S1", s1.Calls)
	fmt.Println("S2", s2.Calls)

}

func TestCountDownDelayBatch(t *testing.T) {
	t.Run("prints 5 to go", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		CountDownDelay(buffer, &CountDownOpSpy{})

		got := buffer.String()
		want := `3
2
1
Go!`
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}

type CountDownOpSpy struct {
	Calls []string
}

func (s *CountDownOpSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountDownOpSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"
