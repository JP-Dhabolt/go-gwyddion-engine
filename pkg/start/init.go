package start

import (
	"github.com/GwyddionGames/go-gwyddion-engine/internal/pkg/private"
	"github.com/GwyddionGames/go-gwyddion-engine/pkg/public"
)

func Init(options public.InitOptions) public.EngineFactory {
	return private.EngineFactory(options)
}
