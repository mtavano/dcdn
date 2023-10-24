package main

import (
	"flag"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/mtavano/dcdn/internal/logger"
	"github.com/mtavano/dcdn/internal/network"
)

func main() {
	log := logger.New("NODE")
	log.Info("Starting")

	// Parse flags
	hostIP := flag.String("host-address", "0.0.0.0", "Default address to run node")
	port := flag.String("node-port", "9000", "Port to enable network connection")
	peerAddresses := flag.String("peer-addresses", "", "Comma separated list of peers to connect with")
	flag.Parse()

	node, err := network.NewNode(&network.NodeConfig{
		IP:   *hostIP,
		Port: *port,
	})
	check(err)

	log.Infof("Runing with address: %s", node.NetworkHost.Addrs())
	log.Infof("with node ID: %s", node.NetworkHost.ID())

	node.Start()
	node.ConnectWithPeers(strings.Split(*peerAddresses, ","))

	defer node.NetworkHost.Close()
	defer node.MdnsService.Close()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGKILL, syscall.SIGINT)
	<-sigCh
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
