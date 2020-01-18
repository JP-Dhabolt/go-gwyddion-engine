package audio

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"os"
	"time"
)

var (
	samplesPerSecond = 441000
	sr = beep.SampleRate(samplesPerSecond)
)

type Engine struct {
	mixer *beep.Mixer
	musicQueue *queue
	effectsQueue *queue
}

func Init() *Engine {
	err := speaker.Init(sr, sr.N(500*time.Millisecond))
	if err != nil {
		panic(err)
	}
	var mixer beep.Mixer
	var musicQueue queue
	var effectsQueue queue
	mixer.Add(&musicQueue, &effectsQueue)
	speaker.Play(&mixer)
	engine := Engine{
		mixer: &mixer,
		musicQueue: &musicQueue,
		effectsQueue: &effectsQueue,
	}
	return &engine
}

func (e *Engine) AddSoundEffect(streamer beep.Streamer) {
	speaker.Lock()
	defer speaker.Unlock()
	e.effectsQueue.Add(streamer)
}

func (e *Engine) AddEffectTone(frequencyInKHz float64, duration time.Duration) {
	speaker.Lock()
	defer speaker.Unlock()
	e.effectsQueue.AddTone(frequencyInKHz, duration)
}

func (e *Engine) AddSoundtrackSong(file *os.File) {
	speaker.Lock()
	defer speaker.Unlock()
	e.musicQueue.AddMp3(file)
}
