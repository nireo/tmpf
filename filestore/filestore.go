package filestore

import "github.com/dgraph-io/badger/v3"

// Filestore handles all operations relating to files.
type Filestore struct {
	Dir      string
	metadata *badger.DB
}

func (fs *Filestore) Cleanup() {
}
