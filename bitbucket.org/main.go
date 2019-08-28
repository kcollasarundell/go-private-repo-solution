package main

import (
	"log"
	"time"
)

type PrivateStuff struct {
}

func NewStuff() *PrivateStuff {
	return &PrivateStuff{}
}

func (p *PrivateStuff) DoStuff() {
	log.Printf("Doing important stuff\n")
	for {
		log.Printf("working...\n")
		time.Sleep(time.Millisecond * 5000)
	}
}

func main() {
	log.Println("Started!")

	z := NewStuff()

	z.DoStuff()
}
