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
		homeDir = os.Getenv("HOME")

		database  = flag.String("database", homeDir+"/.mustekala", "Location of the discovery levelDB")
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
	config.Database = *database
	config.Network = *network

	// There are three components in this application running in parallel
	//
	// * Discovery
	//     Runs the geth p2p server filling a levelDB with discovered nodes
	// * Inquirer
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

	// Discovery
	mustekala.SetupDiscoveryServer(config)
	go mustekala.StartDiscoveryServer()

	// Inquirer
	// TODO

	// API
	// TODO

	// Let the goroutines do their thing
	select {}
}
