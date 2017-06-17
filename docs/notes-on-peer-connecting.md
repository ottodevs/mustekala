# Notes on the Peer Connecting Process of the p2p library

Rough notes taken on the process of connecting and communicating to a p2p node
using the libraries found in `geth`.

## Geth version

```
commit dd06c8584368316c8fb388384b0723d8c7e543f0
Merge: ae40d51 08959bb
Author: Péter Szilágyi <peterke@gmail.com>
Date:   Mon May 29 11:42:48 2017 +0300

    Merge pull request #14523 from karalabe/txpool-cli-flags

    cmd, core, eth: configurable txpool parameters
```

## Preliminaries

* Server starts at `Start()` in `p2p/server.go:344`
  * It invokes `run()` (at `p2p/server.go:417`) as a _goroutine_.
    * `run()` resides at `p2p/server.go:451`. It persists a `for` loop.

## Adding a dial task

* The _adding of new dial tasks_ happens in `scheduleTasks()`, invoked at `p2p/server.go:499` and defined at `p2p/server.go:487`.
  * Inside this function, there is a call to `dialstate.newTasks()`, defined at `p2p/dial.go:141`.
  * The latter function, makes extensive use of `addDial()`, defined at `p2p/dial.go:147` (just inside `newTasks()`).
* Everytime a new task is queued, the function `startTasks()` (in `p2p/server.go:477`) is invoked.
  * If the length of running tasks is lesser than the constant `maxActiveDialTasks`, then the method of the task `Do()` (in `p2p/dial.go:270`) is kicked in.
  * ... And the task added to the `runningTasks` slide.
* `Do()` will deal with the actual dial, defined in `dial()`, in `p2p/dial.go:320`.
* If the actual dial succeeds, we pass to the following stage, `setupConn()` (`p2p/server.go:677`), invoked in `p2p/dial.go:326`.
* The documentation of `setupConn` says

```go
// setupConn runs the handshakes and attempts to add the connection
// as a peer. It returns when the connection has been added as a peer
// or the handshakes have failed.
```

## Setting up the connection after dialing: Encrypted Handshake

* After a dial is successful in `dial()` (`p2p/dial.go:320`), `setupConn()` (`p2p/server.go:677`) is invoked.
* A new object `conn` (`p2p/server.go:190`) is created.
  * `conn` wraps a network connection with information gathered during the two handshakes.
* The function will try an encryption handshake `doEncHandshake()` (`p2p/rplx.go:163`)
  * Which in turn calls `initiatorEncHandshake()` (`p2p/rlpx.go:268`).
* `initiatorEncHandshake()` will package and seal an "`authPacket`" (a `[]byte`) and will send it over the wire.
* When the response has been received, `encHandshake.secrets()` (`p2p/rlpx:227`) is invoked.
  * Getting in response a `secrets` object (`p2p/rlpx.go:195`), which is returned in turn.
* This `secrets` object is delivered to `newRLPXFrameRW()` (`p2p/rlpx.go:561`)
  * Which returns an `rlpxFrameRW` (`p2p/rlpx.go:551`) object, an encapsulation of the `conn` alongside encrypting elements.
* Finally the flow is restored to the function `setupConn()`
  * Two checks are made afterwards:
    * Public ID match between the obtained one after the encrypted handshake, and the node's one in memory.
    * Whether an error happened deliverying the `Server.checkpoint()` (`p2p/server.go:737`)
      * Notifying the `run()` loop (see above) that the node is in a _posthandshake_ status.

## Setting up the connection after dialing: Protocol Handshake

Is time for the protocol handshake!

* This flow is following next to the encryption handshake, described above, inside the function `setupConn()`.
  * The function `doProtoHandshake()` (`p2p/rlpx.go:116`) is invoked at this point in the line (`p2p/server.go:707`).
* According with the in-code documentation, _the protocol handshake is the first authenticated message_.
* Two concurrent operations are performed:
  * `Send()` (`p2p/message.go:92`): Sends the _handshake_ message over the pipe.
  * `readProtocolHandshake()` (`p2p/rlpx.go:133`): Which invokes `rlpxFrameRW.ReadMsg()` (`p2p.rlpx.go:625`),
    * which does the actual processing of the received message.i
    * After parsing this message, `readProtocolHandshake()`, will determine whether the protocol handshake worked or not.
* `doProtoHandshake()`, having found no error, returns _their_ handshake to `setupConn()`.
* Following that, the `id` of the peer is verified, alongside their _capabilities_.
* Finally, the `server.checkpoint()` (`p2p/server.go:737`) is called, and `run()` should be invoking `addPeer()`

## Adding the peer

* Inside `run()` in the `<-srv.addPeer` case (`p2p/server.go:540`), `protoHandshakeChecks()` (`p2p/server.go:587`) is invoked.
  * After checking that this peer has matching protocols, the peer is finally added with `runPeer()` (`p2p/server.go:754`).
* `RunPeer()` calls the goroutine `peer.run()` (`p2p/peer.go:143`), which will be notated below.

## Your connected peer loop

(TODO)

## Talking with your peer: The ethereum sub-protocol

(TODO)
