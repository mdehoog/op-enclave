package da

import (
	"context"
	"encoding/base64"
	"os"
	"path/filepath"

	altda "github.com/ethereum-optimism/optimism/op-alt-da"
)

type filestore struct {
	path string
}

var _ altda.KVStore = &filestore{}

func NewFilestore(path string) altda.KVStore {
	return &filestore{
		path: path,
	}
}

func (s *filestore) Get(ctx context.Context, key []byte) ([]byte, error) {
	return os.ReadFile(filepath.Join(s.path, base64.RawURLEncoding.EncodeToString(key)))
}

func (s *filestore) Put(ctx context.Context, key []byte, value []byte) error {
	return os.WriteFile(filepath.Join(s.path, base64.RawURLEncoding.EncodeToString(key)), value, 0644)
}
