package transient_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/aliworkshop/terra-sdk/store/types"
	"github.com/aliworkshop/terra-sdk/store/v2/transient"
)

var k, v = []byte("hello"), []byte("world")

func TestTransientStore(t *testing.T) {
	tstore := transient.NewStore()
	require.Nil(t, tstore.Get(k))
	tstore.Set(k, v)
	require.Equal(t, v, tstore.Get(k))
	tstore.Commit()
	require.Nil(t, tstore.Get(k))

	emptyCommitID := tstore.LastCommitID()
	require.Equal(t, emptyCommitID.Version, int64(0))
	require.True(t, bytes.Equal(emptyCommitID.Hash, nil))
	require.Equal(t, types.StoreTypeTransient, tstore.GetStoreType())
}
