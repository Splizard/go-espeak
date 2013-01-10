package main

import "github.com/Splizard/go-espeak/espeak"

func main() {
	if err := espeak.Init(); err == -1 {
		return
	}
	espeak.Say("Hello World")
	espeak.Sync()
}
