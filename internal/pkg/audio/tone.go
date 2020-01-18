package audio

import (
	"math"
)

type Tone struct {
	currentSample  int
	frequencyInKHz float64
}

func (t *Tone) Stream(samples [][2]float64) (n int, ok bool) {
	for i := range samples {
		currentFreq := t.frequencyInKHz * 10000
		currentValue := math.Sin(2 * math.Pi * currentFreq * (float64(t.currentSample) / float64(samplesPerSecond)))
		samples[i][0] = currentValue
		samples[i][1] = currentValue
		t.currentSample++
		if t.currentSample > samplesPerSecond {
			t.currentSample = 0
		}
	}
	return len(samples), true
}

func (t *Tone) Err() error {
	return nil
}
