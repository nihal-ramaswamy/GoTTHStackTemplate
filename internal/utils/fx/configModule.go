package fx_utils

import (
	serverconfig "github.com/nihal-ramaswamy/gotta_stack/internal/config/server"
	"go.uber.org/fx"
)

var ConfigModule = fx.Module(
	"Config",
	fx.Provide(serverconfig.Default),
)
