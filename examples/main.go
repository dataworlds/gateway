package main

import (
	"github.com/dataworlds/gateway"
)

func main() {
	proxy := gateway.New()
	proxy.ResourcesFolder = "este"

	proxy.Allow = func() (bool, string) {
		return true, "nose"
	}

	// 	proxy.BeforeSpawn = func(p *player.Player) {
	// 		fmt.Println("test")
	// 	}
	//
	// 	proxy.OnSpawn = func(p *player.Player) {
	// 		fmt.Println("test")
	// 	}

	proxy.Listen(":19132")
}
