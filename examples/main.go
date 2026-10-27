package main

import (
	"github.com/dataworlds/gateway"
	"github.com/dataworlds/gateway/player"
)

func main() {
	proxy := gateway.New()
	proxy.ResourcesFolder = "este"

	proxy.Allow = func() (bool, string) {
		return true, "nose"
	}

	proxy.RegisterDefaultServer("hub", ":19133")

	proxy.OnSpawn("hub", func(p *player.Player) {
	})

	proxy.Listen(":19132")
}
