package start

import (
	"dev.azure.com/gwyddiongames/_git/go-gwyddion-engine/internal/pkg/private"
	"dev.azure.com/gwyddiongames/_git/go-gwyddion-engine/pkg/public"
)

func Init(options public.InitOptions) public.EngineFactory {
	return private.EngineFactory(options)
}
