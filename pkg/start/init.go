package start

import (
	"github.com/JP-Dhabolt/go-gwyddion-engine/internal/pkg/private"
	"github.com/JP-Dhabolt/go-gwyddion-engine/pkg/public"
)

func Init(options public.InitOptions) public.EngineFactory {
	return private.EngineFactory(options)
}
