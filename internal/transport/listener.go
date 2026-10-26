package nethernet

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/df-mc/go-nethernet"
	"github.com/df-mc/go-nethernet/endpoint"
)

func Listen(conn *nethernet.Conn) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
	defer cancel()

	client := endpoint.NewClient()

	slog.Info("establishing connection over NetherNet... send an interrupt signal (Ctrl+C) to abort", "url", "127.0.0.1:19133")
	conn, err := nethernet.Dialer{DisableTrickleICE: true}.DialContext(ctx, "127.0.0.1:19133", client)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	slog.Info("connected",
		"remoteAddr", conn.RemoteAddr(),
		"localAddr", conn.LocalAddr(),
		"latency", conn.Latency(),
	)

	slog.Info("disconnecting in 15 seconds...")
	select {
	case <-ctx.Done():
	case <-conn.Context().Done():
	case <-time.After(time.Second * 15):
	}
}
