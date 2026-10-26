package transport

import (
	"log"
	"net"

	"github.com/df-mc/go-nethernet"
	"github.com/df-mc/go-nethernet/endpoint"
)

func Run(addr string, allow func() (bool, string)) {
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
		go handleConnection(conn)
	}
}

func handleConnection(conn2 net.Conn) {
	client := endpoint.NewClient()

	conn, err := (nethernet.Dialer{
		DisableTrickleICE:       true,
		AllowIdentitylessServer: true,
	}).Dial("http://127.0.0.1:19132", client)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	packet, err := conn.ReadPacket()
	if err != nil {
		log.Printf("Conexión cerrada o error al leer: %v", err)
		return
	}
	log.Printf("Recibí un paquete de %d bytes", len(packet))
}
