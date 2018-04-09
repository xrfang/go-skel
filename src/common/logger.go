package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"
)

var DEBUG_TARGETS []string

func Error(err error) {
	fmt.Fprintln(os.Stderr, trace(err.Error()))
}

func Log(msg string, args ...interface{}) {
	msg = strings.TrimRightFunc(fmt.Sprintf(msg, args...), unicode.IsSpace)
	fmt.Println(msg)
}

func Dbg(msg string, args ...interface{}) {
	var wanted bool
	if len(DEBUG_TARGETS) == 0 {
		wanted = false
	} else if DEBUG_TARGETS[0] == "*" {
		wanted = true
	} else {
		caller := ""
		log := trace("")
		if len(log) > 1 {
			caller = log[1]
		}
		if caller == "" {
			wanted = true
		} else {
			for _, t := range DEBUG_TARGETS {
				if strings.HasSuffix(caller, t) {
					wanted = true
					break
				}
			}
		}
	}
	if wanted {
		Log(msg, args...)
	}
}

func SetDebugTargets(targets string) {
	DEBUG_TARGETS = strings.Split(targets, ",")
}

func Perf(tag string, work func()) {
	start := time.Now()
	Dbg("[EXEC]%s", tag)
	work()
	elapsed := time.Since(start).Seconds()
	Dbg("[DONE]%s (elapsed: %f)", tag, elapsed)
}
