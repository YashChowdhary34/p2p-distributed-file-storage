package p2p

import "errors"

var ErrInvalidHandshake = errors.New("invalid handshake")

// handshakefunc
type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }