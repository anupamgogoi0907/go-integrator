package logger

import "core-esb/pkg/util"

type Logger struct {
}

func (l *Logger) Process() {
	util.ParseApp("data")
}
