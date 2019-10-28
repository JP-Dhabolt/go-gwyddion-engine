package start

import (
	"dev.azure.com/gwyddiongames/_git/go-gwyddion-engine.git/internal/pkg/private"
	"dev.azure.com/gwyddiongames/_git/go-gwyddion-engine.git/pkg/public"
)

func Init(options public.InitOptions) public.EngineFactory {
	return private.EngineFactory(options)
}
