package main

import (
	"sync/atomic"
)

// A simple block manager
type MemBlockManager struct {
	storage   map[uint64][]byte
	currentID uint64
}

func NewMemBlockManager() *MemBlockManager {
	return &MemBlockManager{
		storage:   make(map[uint64][]byte),
		currentID: 0,
	}
}

func (m *MemBlockManager) GetBlock(id uint64) ([]byte, error) {
	res, ok := m.storage[id]
	if !ok || res == nil {
		return nil, ErrNoBlock
	}
	return res, nil
}

func (m *MemBlockManager) SetBlock(id uint64, blk []byte) error {
	res, ok := m.storage[id]
	if !ok || res == nil {
		return ErrNoBlock
	}
	m.storage[id] = blk
	return nil
}

func (m *MemBlockManager) RemoveBlock(id uint64) error {
	res, ok := m.storage[id]
	if !ok || res == nil {
		return ErrNoBlock
	}
	m.storage[id] = nil
	return nil
}

func (m *MemBlockManager) AllocBlock() (uint64, error) {
	res := atomic.AddUint64(&m.currentID, 1)
	m.storage[res] = make([]byte, 0)
	return res, nil
}

func (m *MemBlockManager) Blocksize() uint64 {
	return SizeUnlimited
}
