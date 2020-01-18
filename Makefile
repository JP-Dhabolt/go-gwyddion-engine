.PHONY: audio run-audio lint

audio:
	go build -o audio.exe cmd/audio/main.go

run-audio: audio
	audio.exe

lint:
	go fmt ./...
