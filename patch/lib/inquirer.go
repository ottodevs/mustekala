package mustekala

import (
	"net"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"github.com/ethereum/go-ethereum/params"
)

// SetupInquirer takes the mustekala config struct and delivers
// a ready to use p2p server to perform the discovery of nodes.
func SetupInquirer(config *Config) {
	// Get a new private key
	nodekey, _ := crypto.GenerateKey()

	// Build a new p2pserver
	p2pConfig := &p2p.Config{
		Mustekala:       true, // This flag enables the instantiation of channels to inquire for the statuses of nodes.
		Name:            "Mustekala",
		NodeDatabase:    "", // Use the ethdb memory database for the nodes
		PrivateKey:      nodekey,
		ListenAddr:      ":20000",
		Dialer:          &net.Dialer{KeepAlive: 60},
		MaxPeers:        1000,
		MaxPendingPeers: 1000,
	}

	// Get the desired bootnodes based on network
	setBootNodes(p2pConfig, config.Network)

	// We are ready
	// TODO
	// Add the p2p.protocol we will use
	discoveryServer = &p2p.Server{Config: *p2pConfig}
}

// StartInquirer launches the p2p server
func StartInquirer() {
	if err := discoveryServer.Start(); err != nil {
		log.Error("Unable to setup the discovery server", "err", err)
		os.Exit(1)
	}

	log.Info("INQUIRER STARTED")
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
