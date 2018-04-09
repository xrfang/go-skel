package main

import (
	"flag"
	"fmt"
)

func work() (err error) {
	Dbg("start working")
	defer catch(&err, func() {
		log := trace("something wrong is caught: %v", err)
		fmt.Println(log)
	})
	panic(fmt.Errorf("a deliberate error"))
}

func main() {
	ver := flag.Bool("version", false, "show version info")
	flag.Parse()
	if *ver {
		fmt.Println(verinfo())
		return
	}
	//SetDebugTargets("*")
	SetDebugTargets("work")
	//SetDebugTargets("main,work")
	fmt.Println("This is [main] project.")
	err := work()
	Dbg("work done")
	fmt.Println("===")
	Error(err)
}
