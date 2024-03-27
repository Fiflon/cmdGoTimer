package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fiflon/gt/duration"
	"github.com/fiflon/gt/player"
	"github.com/fiflon/gt/window"
)

func main() {
	// Creates an object of Duration
	d := duration.NewDuration()

	if len(os.Args) != 4 {
		fmt.Fprintln(os.Stderr, "Wrong amount of arguments given. Correct order of arguments is {[number of HOURS] [number of MINUTES] [number of SECONDS]}")
		os.Exit(1) // Zakończ program z kodem błędu
	}

	// Get arguments
	scriptParams := os.Args[1:4]
	d.SetDuration(scriptParams)

	// Waits for a duration that was set
	waitTime := time.Duration(d.Seconds()+d.Minutes()*60+d.Hours()*3600) * time.Second
	fmt.Printf("%d hours, %d minutes, %d seconds - timer starts now! \n", d.Hours(), d.Minutes(), d.Seconds())
	<-time.After(waitTime)
	window.FindAndFocusTerminalWindow()
	fmt.Println("The time is up!")

	// play the sound
	filename := "tanczace_eurydyki.mp3"
	if err := player.PlaySound(filename); err != nil {
		fmt.Fprintf(os.Stderr, "Error while playing sound: %v\n", err)
		os.Exit(1) // Zakończ program z kodem błędu
	}
}
