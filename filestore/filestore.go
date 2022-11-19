package filestore

import (
	"time"

	"github.com/dgraph-io/badger/v3"
	"github.com/gogo/protobuf/proto"
	"github.com/nireo/tmpf/pb"
)

// Filestore handles all operations relating to files.
type Filestore struct {
	Dir      string
	metadata *badger.DB
}

func New(dir string) (*Filestore, error) {
	db, err := badger.Open(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		return nil, err
	}

	return &Filestore{
		Dir:      dir,
		metadata: db,
	}, nil
}

func (fs *Filestore) Add(uuid string, filename string) error {
	meta := &pb.Metadata{
		Filename: filename,
	}

	b, err := proto.Marshal(meta)
	if err != nil {
		return err
	}

	if err = fs.metadata.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte(uuid), b).WithTTL(time.Hour * 24)
		return txn.SetEntry(e)
	}); err != nil {
		return err
	}

	return nil
}

func (fs *Filestore) Get(uuid string) (*pb.Metadata, error) {
	var b []byte
	err := fs.metadata.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(uuid))
		if err != nil {
			return err
		}

		b, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	meta := &pb.Metadata{}
	if err = proto.Unmarshal(b, meta); err != nil {
		return nil, err
	}

	return meta, nil
}
