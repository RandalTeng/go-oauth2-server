package store

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/RandalTeng/oauth2/definition"
)

// NewMemoryTokenStore create a token store instance based on memory map
func NewMemoryTokenStore(cap int) definition.TokenStore {
	return &MemoryTokenStore{
		db:   make(map[string]any),
		lock: &sync.RWMutex{},
		cap:  cap,
	}
}

// MemoryTokenStore token storage based on memory map.
type MemoryTokenStore struct {
	db   map[string]any
	lock *sync.RWMutex
	cap  int
}

// Create and store the new token information
func (ts *MemoryTokenStore) Create(_ context.Context, info definition.TokenInfo) error {
	return ts.set(info.GetAccess(), info)
}

// get key
func (ts *MemoryTokenStore) get(key string) (any, error) {
	ts.lock.RLock()
	defer ts.lock.RUnlock()
	if info, ok := ts.db[key]; ok {
		return info, nil
	} else {
		return nil, errors.New("授权token不存在")
	}
}

// set key
func (ts *MemoryTokenStore) set(key string, info any) error {
	ts.lock.Lock()
	defer ts.lock.Unlock()
	ts.db[key] = info
	return nil
}

// remove key
func (ts *MemoryTokenStore) remove(key string) error {
	ts.lock.Lock()
	defer ts.lock.Unlock()
	delete(ts.db, key)
	return nil
}

// RemoveByCode use the authorization code to delete the token information
func (ts *MemoryTokenStore) RemoveByCode(_ context.Context, code string) error {
	return ts.remove(code)
}

// RemoveByAccess use the access token to delete the token information
func (ts *MemoryTokenStore) RemoveByAccess(_ context.Context, access string) error {
	return ts.remove(access)
}

// RemoveByRefresh use the refresh token to delete the token information
func (ts *MemoryTokenStore) RemoveByRefresh(_ context.Context, refresh string) error {
	return ts.remove(refresh)
}

// GetByCode use the authorization code for token information data
func (ts *MemoryTokenStore) GetByCode(_ context.Context, code string) (definition.TokenInfo, error) {
	value, err := ts.get(code)
	if err != nil {
		return nil, err
	}
	info, ok := value.(definition.TokenInfo)
	if !ok {
		return nil, errors.New(fmt.Sprintf("类型错误: %+v", value))
	} else {
		return info, nil
	}
}

// GetByAccess use the access token for token information data
func (ts *MemoryTokenStore) GetByAccess(_ context.Context, access string) (definition.TokenInfo, error) {
	value, err := ts.get(access)
	if err != nil {
		return nil, err
	}
	info, ok := value.(definition.TokenInfo)
	if !ok {
		return nil, errors.New(fmt.Sprintf("类型错误: %+v", value))
	} else {
		return info, nil
	}
}

// GetByRefresh use the refresh token for token information data
func (ts *MemoryTokenStore) GetByRefresh(_ context.Context, refresh string) (definition.TokenInfo, error) {
	value, err := ts.get(refresh)
	if err != nil {
		return nil, err
	}
	info, ok := value.(definition.TokenInfo)
	if !ok {
		return nil, errors.New(fmt.Sprintf("类型错误: %+v", value))
	} else {
		return info, nil
	}
}
