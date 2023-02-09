package logger

import (
	"fmt"
	"io"
)

type logger struct {
	out io.Writer
}

type Logger interface {
	Log(...interface{})
}

func (lg *logger) Log(args ...interface{}) {
	fmt.Fprint(lg.out, args...)
	fmt.Fprintln(lg.out)
}

// looks like this function should return a *logger rather than Logger
func New(w io.Writer) Logger {
	return &logger{out: w}
}
