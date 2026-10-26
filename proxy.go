package gateway

import (
	"fmt"

	"github.com/dataworlds/gateway/cluster"
	"github.com/dataworlds/gateway/nethernet"
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

func (proxy *proxy) BeforeSpawn(server string, p *player.Player) {

}

func (proxy *proxy) OnSpawn(server string, p *player.Player) {

}

func (proxy *proxy) Listen(addr string) {
	conn, _ := nethernet.Server()
	conn.Context()

	if proxy.ResourcesFolder != "" {
		fmt.Println("loading resourcespacks")
	}
	if proxy.Allow != nil {
		a, b := proxy.Allow()
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
func (proxy *proxy) RegisterDefaultBackend() {
	cluster.RegisterDefaultBackend()

}
