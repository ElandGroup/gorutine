package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/okzk/ticker"
)

//import "sync"

//
func Say(name string) {
	for index := 0; index < 2; index++ {
		runtime.Gosched() //means that the CPU gives the time slice to others
		fmt.Println(name)
	}
}

//var mutex = &sync.Mutex{}

func Routine1() {
	go Say("world")
	Say("hello")
}

func Routine2() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go Say("world")
	Say("hello")
}

//timer
func Routine3() {
	go func() {
		t := time.NewTicker(3 * time.Second) // Every 3 seconds
		for {
			select {
			case <-t.C:
				fmt.Println(time.Now())
			}
		}
		t.Stop()
	}()
}

func Routine4() (newTicker *ticker.Ticker) {
	newTicker = ticker.New(2*time.Second, func(t time.Time) {
		fmt.Println("Tick at", t)
	})
	return
}
