package network

import (
	"fmt"

	"github.com/libp2p/go-libp2p/core/network"
)

func (n *Node) Start() {
	// Set up a stream handler for the protocol on port 9000
	n.logger.Infof("setupping stream")
	n.NetworkHost.SetStreamHandler(Protocol, func(stream network.Stream) {
		// Handle the incoming stream here
		n.logger.Infof("Received incoming connection on port %s", n.port)

		// You can read and write data on the stream as needed
		// ...

		// Close the stream when you're done with it
		err := stream.Close()
		if err != nil {
			fmt.Println("Error closing the stream:", err)
			n.logger.Infof("Error closing the stream:", err)
		}
	})

	// Your other initialization code here...

	// Start listening for incoming connections
	go func() {
		for {
			hosts := n.NetworkHost.Network().Peers()
			if len(hosts) != 0 {
				fmt.Println(hosts)
				break
			}
		}
	}()
}
