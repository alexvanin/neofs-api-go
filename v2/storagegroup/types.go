package storagegroup

import (
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
)

// StorageGroup is a unified structure of StorageGroup
// message from proto definition.
type StorageGroup struct {
	size uint64

	hash []byte

	exp uint64

	members []*refs.ObjectID
}

// GetValidationDataSize of unified storage group structure.
func (s *StorageGroup) GetValidationDataSize() uint64 {
	if s != nil {
		return s.size
	}

	return 0
}

// SetValidationDataSize into unified storage group structure.
func (s *StorageGroup) SetValidationDataSize(v uint64) {
	if s != nil {
		s.size = v
	}
}

// GetValidationHash of unified storage group structure.
func (s *StorageGroup) GetValidationHash() []byte {
	if s != nil {
		return s.hash
	}

	return nil
}

// SetValidationHash into unified storage group structure.
func (s *StorageGroup) SetValidationHash(v []byte) {
	if s != nil {
		s.hash = v
	}
}

// GetExpirationEpoch of unified storage group structure.
func (s *StorageGroup) GetExpirationEpoch() uint64 {
	if s != nil {
		return s.exp
	}

	return 0
}

// SetExpirationEpoch into unified storage group structure.
func (s *StorageGroup) SetExpirationEpoch(v uint64) {
	if s != nil {
		s.exp = v
	}
}

// GetMembers of unified storage group structure. Members are objects of
// storage group.
func (s *StorageGroup) GetMembers() []*refs.ObjectID {
	if s != nil {
		return s.members
	}

	return nil
}

// SetMembers into unified storage group structure. Members are objects of
// storage group.
func (s *StorageGroup) SetMembers(v []*refs.ObjectID) {
	if s != nil {
		s.members = v
	}
}
