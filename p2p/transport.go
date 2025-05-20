package p2p

// peer is an interface that represents the remote node
type Peer interface {
	close() error
}

// transport is anything that handles the oommunication between the nodes in the network. This can be of the form TCP, UDP, websockets...
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}