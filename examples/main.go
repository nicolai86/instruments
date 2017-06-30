package main

import (
	"fmt"
	"time"

	"github.com/nicolai86/instruments"
)

func main() {
	fmt.Printf("Hello world\n")
	time.Sleep(10 * time.Second)

	go func() {
		<-time.After(1 * time.Second)
		instruments.Signpost(42)
	}()

	region := instruments.StartWithArguments(43, 0, 0, 0, instruments.ColorPurple)

	time.Sleep(5 * time.Second)

	region.End()
}
