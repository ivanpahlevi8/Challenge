package configs

import (
	"database/sql"

	"github.com/alexedwards/scs/v2"
)

// create config object
type Config struct {
	DB           *sql.DB
	Session      *scs.SessionManager
	InProduction bool
}
