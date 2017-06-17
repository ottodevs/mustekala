package main

import (
	"flag"
	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/mustekala"
)

func main() {
	// Parameters and Initial Setup
	var (
		network   = flag.String("network", "mainnet", "bootnodes type")
		verbosity = flag.Int("verbosity", int(log.LvlInfo), "log verbosity (0-9)")
		vmodule   = flag.String("vmodule", "", "log verbosity pattern")
	)
	flag.Parse()

	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(log.Lvl(*verbosity))
	glogger.Vmodule(*vmodule)
	log.Root().SetHandler(glogger)

	config := &mustekala.Config{}
	config.Network = *network

	// There are two components in this application running in parallel
	//
	// * Inquirer
	//     Runs the geth p2p server filling a levelDB with discovered nodes
	//     Tries to connect to the discovered servers using the p2p library,
	//     classifying them into the following categories:
	//       - TCP Dial failed
	//       - Encryption Handshake failed
	//       - Protocol Handshake failed
	//       - Get Status failed
	//       - Get Status succeed
	// * API
	//    HTTP JSON API. Serves the following requests:
	//      - /discovered-nodes
	//          Information of each discovered enode, plus latest connection attempt data
	//      - /network-status
	//          Summary of the network, grouping the nodes by statuses

	// Memory DB
	// TODO

	// Inquirer
	mustekala.SetupInquirer(config)
	go mustekala.StartInquirer()

	// API
	// TODO

	// Let the goroutines do their thing
	select {}
}
