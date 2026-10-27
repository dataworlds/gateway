package cluster

import (
	"errors"
	"fmt"
	"strings"
)

func RegisterServer(name, addr string, makeDefault bool) error {
	name, addr = strings.TrimSpace(name), strings.TrimSpace(addr)
	if name == "" {
		return errors.New("gateway: server name cannot be empty")
	}
	if addr == "" {
		return fmt.Errorf("gateway: address for server %q cannot be empty", name)
	}
	if !strings.HasPrefix(addr, "http://") && !strings.HasPrefix(addr, "https://") {
		return fmt.Errorf("gateway: server %q address must start with http:// or https://", name)
	}

	return nil
}
