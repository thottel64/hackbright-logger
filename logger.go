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

func New(w io.Writer) *logger {
	return &logger{out: w}
}
