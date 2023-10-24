package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/libp2p/go-libp2p"
)

func main() {
	//ctx := context.Background()

	host, err := libp2p.New(nil)
	if err != nil {
		panic(err)
	}
	defer host.Close()

	fmt.Println(host.Addrs())
	fmt.Println("ID:", host.ID())

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGKILL, syscall.SIGINT)
	<-sigCh
}
