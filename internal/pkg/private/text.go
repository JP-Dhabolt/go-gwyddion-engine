package private

import (
	"fmt"
	"github.com/4ydx/gltext"
	"github.com/4ydx/gltext/v4.1"
	"github.com/go-gl/mathgl/mgl32"
	"golang.org/x/image/math/fixed"
	"os"
	"path/filepath"
)

// loadFont loads the specified font at the given scale.
func loadFont(file string, scale int32) (*v41.Font, error) {

	//absPath, _ := filepath.Abs(file)
	//fd, err := os.Open(absPath)
	//if err != nil {
	//	return nil, err
	//}
	//defer fd.Close()
	//
	//runeRanges := make(gltext.RuneRanges, 0)
	//runeScale := fixed.Int26_6(scale)
	//runesPerRow := fixed.Int26_6(128)
	//
	//config, err := gltext.NewTruetypeFontConfig(fd, runeScale, runeRanges, runesPerRow)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return v41.NewFont(config)

	fontConfigDir := "fontconfigs"
	var fontConfigName = fmt.Sprintf("%s%s%d", file, "_", scale)
	var font *v41.Font
	config, err := gltext.LoadTruetypeFontConfig(fontConfigDir, fontConfigName)
	if err == nil {
		font, err = v41.NewFont(config)
		if err != nil {
			return nil, err
		}
		fmt.Println("Font loaded from disk...")
		return font, nil
	} else {
		relativeFilePath := filepath.Join(".", file + ".ttf")
		absPath, _ := filepath.Abs(relativeFilePath)
		fd, err := os.Open(absPath)
		if err != nil {
			return nil, err
		}
		defer func() {
			_ = fd.Close()
		}()

		// Japanese character ranges
		// http://www.rikai.com/library/kanjitables/kanji_codes.unicode.shtml
		runeRanges := make(gltext.RuneRanges, 0)
		runeRanges = append(runeRanges, gltext.RuneRange{Low: 32, High: 128})
		runeRanges = append(runeRanges, gltext.RuneRange{Low: 0x3000, High: 0x3030})
		runeRanges = append(runeRanges, gltext.RuneRange{Low: 0x3040, High: 0x309f})
		runeRanges = append(runeRanges, gltext.RuneRange{Low: 0x30a0, High: 0x30ff})
		runeRanges = append(runeRanges, gltext.RuneRange{Low: 0x4e00, High: 0x9faf})
		runeRanges = append(runeRanges, gltext.RuneRange{Low: 0xff00, High: 0xffef})

		runeScale := fixed.Int26_6(24*scale)
		runesPerRow := fixed.Int26_6(128)
		adjustedHeight := fixed.Int26_6(0)
		config, err = gltext.NewTruetypeFontConfig(fd, runeScale, runeRanges, runesPerRow, adjustedHeight)
		if err != nil {
			panic(err)
		}
		config.Name = fontConfigName

		err = config.Save(fontConfigDir, fontConfigName)
		if err != nil {
			panic(err)
		}
		return v41.NewFont(config)
	}
}

// drawString draws the same string for each loaded font.
func (engine *gameEngine) drawString(x, y float32, str string) error {
	if engine.font != nil {
		width, height := engine.window.GetSize()
		engine.font.ResizeWindow(float32(width), float32(height))

		scaleMin, scaleMax := float32(1.0), float32(1.1)
		text := v41.NewText(engine.font, scaleMin, scaleMax)
		text.SetString(str)
		text.SetColor(mgl32.Vec3{1, 1, 1})
		//text.FadeOutPerFrame = 0.01
		//sw, sh := engine.font.Metrics(str)
		//engine.setColor(public.Color{Red: 0.1, Green: 0.1, Blue: 0.1, Alpha: 0.7})
		//gl.Rectf(x, y, x+float32(sw), y+float32(sh))

		// Render the string.
		//engine.setColor(public.Color{Red: 0, Green: 0, Blue: 0, Alpha: 1})

		//gl.ClearColor(0.4, 0.4, 0.4, 0.0)
		//
		//gl.Clear(gl.COLOR_BUFFER_BIT)

		// Position can be set freely
		text.SetPosition(mgl32.Vec2{0, float32(450)})
		text.Draw()

	}
	return nil
}