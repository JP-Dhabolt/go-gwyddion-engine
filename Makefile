.PHONY: audio run-audio

audio:
	go build -o audio.exe cmd/audio/main.go

run-audio: audio
	audio.exe
