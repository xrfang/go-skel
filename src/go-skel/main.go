package main

import (
	"flag"
	"fmt"
)

func work() (err error) {
	defer catch(&err, func() {
		log := trace("something wrong is caught: %v", err)
		for _, l := range log {
			fmt.Println(l)
		}
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
	fmt.Println("This is [main] project.")
	err := work()
	fmt.Printf("ERR=%v\n", err)
}
