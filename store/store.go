package store

import (
	dbm "github.com/tendermint/tm-db"

	"github.com/aliworkshop/terra-sdk/store/cache"
	"github.com/aliworkshop/terra-sdk/store/rootmulti"
	"github.com/aliworkshop/terra-sdk/store/types"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	return rootmulti.NewStore(db)
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	return cache.NewCommitKVStoreCacheManager(cache.DefaultCommitKVStoreCacheSize)
}
