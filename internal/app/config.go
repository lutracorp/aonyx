package app

import (
	"github.com/lutracorp/lutraauth/internal/pkg/database"
	"github.com/lutracorp/lutraauth/internal/pkg/server"
)

// Config specifies LutraAuth configuration.
type Config struct {
	Server   *server.Config   `hcl:"server,block"`   // Server specifies server configuration.
	Database *database.Config `hcl:"database,block"` // Database specifies database connection configuration.
}
