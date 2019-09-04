package resolvers

import (
	"github.com/korolr/go-sh/db"
)

// Resolvers including query and mutation
type Resolvers struct {
	*db.DB
}
