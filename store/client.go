package store

import (
	"context"
	"errors"
	"sync"

	"github.com/RandalTeng/go-oauth2-server/definition"
)

// NewClientStore create client store
func NewClientStore() *ClientStore {
	return &ClientStore{
		data: make(map[string]definition.ClientInfo),
	}
}

// ClientStore client information store
type ClientStore struct {
	sync.RWMutex
	data map[string]definition.ClientInfo
}

// GetByID according to the ID for the client information
func (cs *ClientStore) GetByID(_ context.Context, id string) (definition.ClientInfo, error) {
	cs.RLock()
	defer cs.RUnlock()

	if c, ok := cs.data[id]; ok {
		return c, nil
	}
	return nil, errors.New("not found")
}

// Set client information
func (cs *ClientStore) Set(id string, cli definition.ClientInfo) (err error) {
	cs.Lock()
	defer cs.Unlock()

	cs.data[id] = cli
	return
}
