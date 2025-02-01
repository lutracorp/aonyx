package app

import (
	"github.com/lutracorp/aonyx/internal/pkg/database"
	"github.com/lutracorp/aonyx/internal/pkg/server"
)

// Config specifies Aonyx configuration.
type Config struct {
	Server   *server.Config   `hcl:"server,block"`   // Server specifies server configuration.
	Database *database.Config `hcl:"database,block"` // Database specifies database connection configuration.
}
