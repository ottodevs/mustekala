# Notes on the Peer Connecting Process of the p2p library

Rough notes taken on the process of connecting to a p2p node.

## Geth version

```
commit dd06c8584368316c8fb388384b0723d8c7e543f0
Merge: ae40d51 08959bb
Author: Péter Szilágyi <peterke@gmail.com>
Date:   Mon May 29 11:42:48 2017 +0300

    Merge pull request #14523 from karalabe/txpool-cli-flags

    cmd, core, eth: configurable txpool parameters
```

## Preliminars

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

(TODO)
