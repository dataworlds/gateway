package nethernet

import (
	"log"

	"github.com/df-mc/go-nethernet"
	"github.com/df-mc/go-nethernet/endpoint"
)

func Dial() {
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
