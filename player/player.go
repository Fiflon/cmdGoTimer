package player

import (
	"math/rand"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
)

type Noise struct{}

func (no Noise) Stream(samples [][2]float64) (n int, ok bool) {
	for i := range samples {
		samples[i][0] = (rand.Float64()*2 - 1) * 0.1
		samples[i][1] = (rand.Float64()*2 - 1) * 0.1
	}
	return len(samples), true
}
func (no Noise) Err() error {
	return nil
}

func PlaySound() error {
	sr := beep.SampleRate(44100)

	speaker.Init(sr, sr.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(beep.Take(sr.N(5*time.Second), Noise{}), beep.Callback(func() {
		done <- true
	})))
	<-done
	return nil
}
