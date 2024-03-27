package duration

import (
	"log"
	"strconv"
)

type Duration struct {
	seconds int
	minutes int
	hours   int
}

func ParseTime(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("Wrong arguments given. Correct order of arguments is {[number of HOURS] [number of MINUTES] [number of SECONDS]}")
	}

	return v
}

func NewDuration() Duration {
	return Duration{seconds: 1, minutes: 0, hours: 0}
}

func (d *Duration) SetDuration(params []string) {
	d.hours = ParseTime(params[0])
	d.minutes = ParseTime(params[1])
	d.seconds = ParseTime(params[2])
}
func (d *Duration) Hours() int {
	return d.hours
}

func (d *Duration) Minutes() int {
	return d.minutes
}

func (d *Duration) Seconds() int {
	return d.seconds
}
