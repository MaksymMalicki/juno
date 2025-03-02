package feedergatewaysync

import (
	"encoding/json"
	"log"

	"github.com/cockroachdb/pebble"
)

type Storage struct {
	Db *pebble.DB
}

func NewStorage(dbPath string) *Storage {
	db, err := pebble.Open(dbPath, nil)
	if err != nil {
		panic(err)
	}
	return &Storage{
		Db: db,
	}
}

func (s *Storage) Close() {
	s.Db.Close()
}

func (s *Storage) Get(key []byte) (*Block, error) {
	value, closer, err := s.Db.Get(key)
	if err != nil {
		return nil, err
	}
	defer closer.Close()

	var block Block
	err = json.Unmarshal(value, &block)
	if err != nil {
		return nil, err
	}

	return &block, nil
}

func (s *Storage) Set(key []byte, block *Block) error {
	value, err := json.Marshal(block)
	if err != nil {
		return err
	}
	log.Printf("Block written to DB")
	return s.Db.Set(key, value, pebble.Sync)
}
