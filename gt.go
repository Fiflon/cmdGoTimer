package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fiflon/gt/duration"
	"github.com/fiflon/gt/player"
	"github.com/fiflon/gt/window"
)

func main() {
	// Creates an object of Duration
	d := duration.NewDuration()

	// Get arguments
	scriptParams := os.Args[1:4]
	d.SetDuration(scriptParams)

	// Waits for a duration that was set
	waitTime := time.Duration(d.Seconds()+d.Minutes()*60+d.Hours()*3600) * time.Second
	fmt.Printf("%d hours, %d minutes, %d seconds - timer starts now! \n", d.Hours(), d.Minutes(), d.Seconds())
	time.After(waitTime)
	window.FindAndFocusTerminalWindow()
	fmt.Println("The time is up!")

	// play the sound
	filename := "tanczace_eurydyki.mp3"
	if err := player.PlaySound(filename); err != nil {
		log.Fatalf("Error while playing sound: %v", err)
	}
}
