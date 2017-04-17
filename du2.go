package main

import (
	"flag"
	"time"
)

var verbose = flag.Bool("b", false, "show verbose program message")

func main() {
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select{
		case size, ok := <= fileSizes:
			if !ok{
				break loop
			}
		}
	}
}
