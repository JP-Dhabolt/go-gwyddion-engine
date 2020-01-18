package main

import (
	"dev.azure.com/gwyddiongames/_git/go-gwyddion-engine.git/internal/pkg/audio"
	"time"
)

func main() {
	engine := audio.Init()
	engine.AddEffectTone(0.5, 1 * time.Hour)
	engine.AddEffectTone(0.8, 750 * time.Millisecond)
	//file, err := os.Open("E:\\SteamLibrary\\steamapps\\common\\Audioshield\\Audioshield_Data\\StreamingAssets\\SampleSongs\\06 364 Days (Album).mp3")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//engine.AddSoundtrackSong(file)
	select {}
}
