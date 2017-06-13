# Mustekala

Ethereum Network Status PoC

## Overview

Mustekala is a patch of `go-ethereum`, to assist the client on obtaining
information about the ethereum network.

### What information can I get from musteka.la

* List of discovered nodes
* Status of connection of these nodes
  * TCP Dial failed
  * Encryption Handshake failed
  * Protocol Handshake failed
  * Get Status failed
  * Get Status succeed

Tipically, when you reach `Get Status succeed`, is when you can actually
connect to a node, and synchronize from it.

### Feedback

You can write [an issue](https://github.com/ConsenSys/mustekala/issues) or,
of course, fork this repository and do a [pull request](https://github.com/ConsenSys/mustekala/pulls).

### Installing the patch

You need to have `go-ethereum` in the directory

```
$GOPATH/src/github.com/ethereum/go-ethereum
```

And run the command

```
./install_patch
```

This will copy the contents of the directory `patch` into the new directories `/mustekala`
and `/cmd/mustekala` in `go-ethereum`, and build your `mustekala` executable.

### Running mustekala

Run `mustekala`

```
$GOPATH/src/github.com/ethereum/go-ethereumbuild/bin/mustekala
```

### Options

* `database`: Path of your levelDB discovery database. Default is `$HOME/.mustekala`.
* `network`: Where to get your bootnodes from params. Default is `mainnet`. 
* `verbosity`: log verbosity (0-9).
* `vmodule`: log verbosity pattern.

### Getting information from `mustekala`

(TODO)
(TODO: API)
