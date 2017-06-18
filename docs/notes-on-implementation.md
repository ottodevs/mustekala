# Notes on Implementation

The most important rule is to use the less amount of code possible, relying on
the `go-ethereum` libraries and modifying, when possible, the host code in
reporting points, avoiding to change the incumbent logic.

## Notes on peer connecting

[This document](nodes-on-peer-connecting.md) contains information related function,
files and line numbers where the important logic points on p2p connection happen.

## Statuses to measure

* Node Discovered
* TCP Dial failed
* Encryption Handshake failed
* Protocol Handshake failed
* Get Status failed
* Get Status succeed

## Mustekala's memory database

Corresponds to a simple private map. It is not designed to be directly accesible,
encouraging the use to invoke its API instead.

## Measuring

### Preliminaries

* The _main_ function of _mustekala_ will set up a `p2p.server.Config` flag called `Mustekala` to true.
  * This flag will enable the instantiation of different channels to be fed by hooks inside the `p2p` and `p2p/discover` libraries.

### Node Discovered

* A hook is added in the function `updateNode()` (`p2p/discovery/database.go:194`).
  * It will send the node information to the `mustekalaDiscoveredCh` channel.
  * A loop in the `mustekala` library will feed from this channel and modify the database accordingly.

### TCP Dial failed

(TODO)

### Encryption Handshake failed

(TODO)

### Protocol Handshake failed

(TODO)

### Get Status failed

(TODO)

### Get Status succeed

(TODO)
