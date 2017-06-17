# Notes on Implementation

The most important rule is to use the less amount of code possible, relying on
the `go-ethereum` libraries and modifying, when possible, the host code in
reporting points, avoiding to change the incumbent logic.

## Notes on peer connecting

[This document](nodes-on-peer-connecting.md) contains information related function,
files and line numbers where the important logic points on p2p connection happen.

## Statuses to measure

* TCP Dial failed
* Encryption Handshake failed
* Protocol Handshake failed
* Get Status failed
* Get Status succeed

## Mustekala's memory database

Corresponds to a simple private map. It is not designed to be directly accesible,
encouraging the use to invoke its API instead.

## Measuring

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
