package mustekala

import (
	"github.com/ethereum/go-ethereum/p2p"

	"github.com/syndtr/goleveldb/leveldb"
)

// Package variables
var (
	discoveryServer *p2p.Server
	levelDBInstance *leveldb.DB
)

type Config struct {
	Database string
	Network  string

	// Keep the following variables private and access them
	// with setters and getters, should you need to
}
