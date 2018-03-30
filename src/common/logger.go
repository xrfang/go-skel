package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"
)

var DEBUG_MODE bool

func Error(err error) {
	fmt.Fprintln(os.Stderr, trace(err.Error()))
}

func Log(msg string, args ...interface{}) {
	msg = strings.TrimRightFunc(fmt.Sprintf(msg, args...), unicode.IsSpace)
	fmt.Println(msg)
}

func Dbg(msg string, args ...interface{}) {
	if DEBUG_MODE {
		Log(msg, args...)
	}
}

func Perf(tag string, work func()) {
	start := time.Now()
	Dbg("[EXEC]%s", tag)
	work()
	elapsed := time.Since(start).Seconds()
	Dbg("[DONE]%s (elapsed: %f)", tag, elapsed)
}
