package audio

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"os"
	"time"
)

type queue struct {
	streamers []beep.Streamer
}

func (q *queue) AddTone(frequency float64, duration time.Duration) {
	sound := beep.Take(sr.N(duration), &Tone{frequencyInKHz: frequency})
	q.Add(sound)
}

func (q *queue) Add(streamers ...beep.Streamer) {
	q.streamers = append(q.streamers, streamers...)
}

func (q *queue) AddMp3(file *os.File) {
	streamer, format, err := mp3.Decode(file)
	if err != nil {
		panic(err) // TODO: Something better than panic here
	}
	resampled := beep.Resample(4, format.SampleRate, sr, streamer)
	q.Add(resampled)
}

func (q *queue) Stream(samples [][2]float64) (n int, ok bool) {
	filled := 0
	for filled < len(samples) {
		if len(q.streamers) == 0 {
			for i := range samples[filled:] {
				samples[i][0] = 0
				samples[i][1] = 0
			}
			break
		}

		n, ok := q.streamers[0].Stream(samples[filled:])
		if !ok {
			q.streamers = q.streamers[1:]
		}
		filled += n
	}
	return len(samples), true
}

func (q *queue) Err() error {
	return nil
}
