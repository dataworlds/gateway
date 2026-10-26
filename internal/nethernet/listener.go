package nethernet

import (
	"log"

	"github.com/df-mc/go-nethernet"
	"github.com/df-mc/go-nethernet/endpoint"
)

func Listen(addr string) {
	signal, err := endpoint.Serve(addr)
	if err != nil {
		log.Fatal(err)
	}
	defer signal.Close()

	listener, err := (nethernet.ListenConfig{
		DisableTrickleICE: signal.DisableTrickleICE(),
	}).Listen(signal)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}

		_ = conn
	}
}
