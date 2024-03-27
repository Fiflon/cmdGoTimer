package player

import (
	"math/rand"
	"os"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
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

func PlaySound(filename string) error {
	sr := beep.SampleRate(44100)
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		return err
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan struct{})
	speaker.Play(beep.Seq(beep.Take(sr.N(8*time.Second), Noise{}), beep.Callback(func() {
		close(done)
	})))

	<-done
	return nil
}
