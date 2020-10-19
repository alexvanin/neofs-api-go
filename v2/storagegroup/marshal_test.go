package storagegroup_test

import (
	"testing"

	goproto "github.com/golang/protobuf/proto"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/storagegroup"
	grpc "github.com/nspcc-dev/neofs-api-go/v2/storagegroup/grpc"
	"github.com/stretchr/testify/require"
)

func TestStorageGroup_StableMarshal(t *testing.T) {
	ownerID1 := new(refs.ObjectID)
	ownerID1.SetValue([]byte("Object ID 1"))

	ownerID2 := new(refs.ObjectID)
	ownerID2.SetValue([]byte("Object ID 2"))

	storageGroupFrom := new(storagegroup.StorageGroup)
	transport := new(grpc.StorageGroup)

	t.Run("non empty", func(t *testing.T) {
		storageGroupFrom.SetValidationDataSize(300)
		storageGroupFrom.SetValidationHash(generateChecksum("Homomorphic hash"))
		storageGroupFrom.SetExpirationEpoch(100)
		storageGroupFrom.SetMembers([]*refs.ObjectID{ownerID1, ownerID2})

		wire, err := storageGroupFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = goproto.Unmarshal(wire, transport)
		require.NoError(t, err)

		storageGroupTo := storagegroup.StorageGroupFromGRPCMessage(transport)
		require.Equal(t, storageGroupFrom, storageGroupTo)
	})
}

func generateChecksum(data string) *refs.Checksum {
	checksum := new(refs.Checksum)
	checksum.SetType(refs.TillichZemor)
	checksum.SetSum([]byte(data))

	return checksum
}
