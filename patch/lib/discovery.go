package mustekala

import (
	"net"
	"reflect"
	"unsafe"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"github.com/ethereum/go-ethereum/params"

	"github.com/syndtr/goleveldb/leveldb"
)

// SetupDiscoveryServer takes the mustekala config struct and delivers
// a ready to use p2p server to perform the discovery of nodes.
func SetupDiscoveryServer(config *Config) {
	// Get a new private key
	nodekey, _ := crypto.GenerateKey()

	// Build a new p2pserver
	p2pConfig := &p2p.Config{
		Name:            "Mustekala",
		NodeDatabase:    config.Database,
		PrivateKey:      nodekey,
		ListenAddr:      ":20000",
		Dialer:          &net.Dialer{KeepAlive: 60},
		MaxPeers:        1000,
		MaxPendingPeers: 1000,
	}

	// Get the desired bootnodes based on network
	setBootNodes(p2pConfig, config.Network)

	// We are ready
	discoveryServer = &p2p.Server{Config: *p2pConfig}
}

// StartDiscoveryServer launches the p2p server and gets the pointer
// of the used levelDB to enable other processes to get the retrieved data
func StartDiscoveryServer() {
	if err := discoveryServer.Start(); err != nil {
		log.Error("Unable to setup the discovery server", "err", err)
	}

	log.Info("Discovery Server Started")

	// Use the hack to get the pointer (and thread) of that levelDB instance
	// and be able to access to the discovered nodes
	levelDBInstance = getLevelDBInstance()

	// Block here to let the server run
	select {}
}

// Based on the network given, it will get the bootnodes configured
// in params/bootnodes.go
// TODO
// Make possible to have an option different to mainnet.
func setBootNodes(cfg *p2p.Config, network string) {
	// TODO
	// Switch over network values
	urls := params.MainnetBootnodes

	cfg.BootstrapNodes = make([]*discover.Node, 0, len(urls))
	for _, url := range urls {
		node, err := discover.ParseNode(url)
		if err != nil {
			log.Error("Bootstrap URL invalid", "enode", url, "err", err)
			continue
		}
		cfg.BootstrapNodes = append(cfg.BootstrapNodes, node)
	}
}

// This is a hack I wouldn't introduce to my parents!
// Does the job anyways, as we are able to get the pointer of
// the levelDB this discovery server is using, without the need to
// patch this component of the p2p library.
func getLevelDBInstance() *leveldb.DB {
	ntab_reflect_value := reflect.
		ValueOf(discoveryServer).
		Elem().
		FieldByName("ntab").
		Elem()

	lvl_reflect_value := reflect.
		ValueOf((*discover.Table)(unsafe.Pointer(ntab_reflect_value.Pointer()))).
		Elem().
		FieldByName("db").
		Elem().
		FieldByName("lvl")

	return (*leveldb.DB)(unsafe.Pointer(lvl_reflect_value.Pointer()))
}
