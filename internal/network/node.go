package network

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
	"github.com/mtavano/dcdn/internal/logger"
	"github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
)

const (
	// public const
	Protocol  = "/protocol-alpha/1.0.0"
	Namespace = "NETWORK"

	//private const
	namespace = "PROTOCOL-ALPA-PEERS"
)

type Node struct {
	NetworkHost host.Host
	MdnsService mdns.Service

	connectedPeers map[string]*peer.AddrInfo
	logger         *logger.Logger
	port           string
}

func NewNode(config *NodeConfig) (*Node, error) {
	log := logger.New(Namespace)

	listenAddress := fmt.Sprintf("/ip4/%s/tcp/%s", config.IP, config.Port)
	address := libp2p.ListenAddrStrings(listenAddress)
	host, err := libp2p.New(address)
	if err != nil {
		return nil, errors.Wrap(err, "network: NewNode libp2p.New error")
	}

	mdnsService := mdns.NewMdnsService(
		host,
		namespace,
		&discoveryveryNotifee{logger: log},
	)

	return &Node{
		NetworkHost:    host,
		MdnsService:    mdnsService,
		connectedPeers: make(map[string]*peer.AddrInfo),
		logger:         log,
		port:           config.Port,
	}, nil
}

func (n *Node) ConnectWithPeers(peersAddresses []string) error {
	for _, peerAddr := range peersAddresses {
		peerMultiAddr, err := multiaddr.NewMultiaddr(peerAddr)
		if err != nil {
			err = errors.Wrap(err, "network: Node.ConnectWithPeers multiaddr.NewMultiaddr error")
			n.logger.Infof("%s  with peer %s", err.Error(), peerAddr)
			return err
		}

		peerAddrInfo, err := peer.AddrInfoFromP2pAddr(peerMultiAddr)
		if err != nil {
			err = errors.Wrap(err, "network: Node.ConnectWithPeers peer.AddrInfoFromP2pAddr error")
			n.logger.Infof("%s  with peer %s", err.Error(), peerAddr)
			return err
		}

		err = n.NetworkHost.Connect(context.Background(), *peerAddrInfo)
		if err != nil {
			err = errors.Wrap(err, "network: Node.ConnectWithPeers n.NetworkHost.Connect error")
			n.logger.Infof("%s  with peer %s", err.Error(), peerAddr)
			return err
		}

		n.connectedPeers[peerAddrInfo.String()] = peerAddrInfo
	}

	return nil
}

type discoveryveryNotifee struct {
	logger *logger.Logger
}

func (dn *discoveryveryNotifee) HandlePeerFound(peerInfo peer.AddrInfo) {
	// TODO: add logic here to store info about peers connected
	dn.logger.Infof("Found peer %s", peerInfo.String())
}
