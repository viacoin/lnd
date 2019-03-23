routing
=======

[![Build Status](http://img.shields.io/travis/viacoin/lnd.svg)](https://travis-ci.org/viacoin/lnd) 
[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/viacoin/lnd/blob/master/LICENSE)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/viacoin/lnd/routing)

The routing package implements authentication+validation of channel
announcements, pruning of the channel graph, path finding within the network,
sending outgoing payments into the network and synchronizing new peers to our
channel graph state.

## Installation and Updating

```bash
$ go get -u github.com/viacoin/lnd/routing
```
