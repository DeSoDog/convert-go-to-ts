package types

import (
	"github.com/dgraph-io/badger/v3"
	"github.com/go-pg/pg/v10"
)
 
type Notifier struct {
	coreChain *Blockchain
	postgres  *Postgres

	// Shortcut to postgres.db
	db *pg.DB

	// Shortcut to coreChain.db
	badger *badger.DB
} 