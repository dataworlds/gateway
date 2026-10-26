package gateway

import (
	"fmt"

	"github.com/dataworlds/gateway/internal/transport"
	"github.com/dataworlds/gateway/player"
)

type proxy struct {
	ResourcesFolder string
	Function        func()
	Allow           func() (bool, string)
}

func New() *proxy {
	return &proxy{}
}

func (proxy *proxy) RegisterProxy(addr string) {

}

func (proxy *proxy) RegisterCommand() {

}

func (proxy *proxy) BeforeSpawn(server string, p *player.Player) {

}

func (proxy *proxy) OnSpawn(server string, p *player.Player) {

}

func (proxy *proxy) Listen(addr string) {
	transport.Run(addr)
	// conn, _ := nethernet.Listen(addr)
	// conn.Context()

	// 	remoteAddr := getServer()
	//
	// 	nethernet.Dial(remoteAddr)
	fmt.Println("test")

	if proxy.ResourcesFolder != "" {
		fmt.Println("loading resourcespacks")
	}
	if proxy.Allow != nil {
		a, b := proxy.Allow()
		_ = a
		_ = b
	}
	// if proxy.BeforeSpawn != nil {
	// 	proxy.BeforeSpawn(&player.Player{})
	// }
	// if proxy.OnSpawn != nil {
	// 	proxy.OnSpawn(&player.Player{})
	// }
}

// backend
// proxy
// cmd
