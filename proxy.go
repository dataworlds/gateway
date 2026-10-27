package gateway

import (
	"github.com/dataworlds/gateway/internal/cluster"
	"github.com/dataworlds/gateway/internal/cmd"
	"github.com/dataworlds/gateway/internal/transport"
	"github.com/dataworlds/gateway/player"
)

type proxy struct {
	Allow           func() (bool, string)
	ResourcesFolder string
}

func New() *proxy {
	return &proxy{}
}

func (proxy *proxy) RegisterProxy(addr string) {
	cluster.RegisterProxy(addr)
}

func (proxy *proxy) RegisterDefaultServer(name, addr string) {
	cluster.RegisterServer(name, addr, true)
}

func (proxy *proxy) RegisterServer(name, addr string) {
	cluster.RegisterServer(name, addr, false)
}

func (proxy *proxy) RegisterCommand() {
	cmd.Register()
}

func (proxy *proxy) Function(func()) {

}

func (proxy *proxy) BeforeSpawn(server string, p *player.Player) {

}

func (proxy *proxy) OnSpawn(server string, p *player.Player) {

}

func (proxy *proxy) Listen(addr string) {
	if proxy.ResourcesFolder != "" {
		// read resources
	}

	transport.Run(addr, proxy.Allow)
}
