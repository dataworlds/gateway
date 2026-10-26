package nethernet

import (
	"errors"
	"log/slog"
	"net"
	"net/http"

	"github.com/df-mc/go-nethernet"
	"github.com/df-mc/go-nethernet/endpoint"
)

func Dial() (*nethernet.Conn, error) {
	server := endpoint.NewHandler()

	l, err := nethernet.ListenConfig{
		AllowAnonymous:    true,
		DisableTrickleICE: true,
	}.Listen(server)
	if err != nil {
		return nil, err
	}

	tcp, err := net.Listen("tcp", ":19132")
	if err != nil {
		_ = l.Close()
		return nil, err
	}

	srv := &http.Server{
		Addr:    ":19132",
		Handler: server,
	}

	go func() {
		if err := srv.Serve(tcp); err != nil &&
			!errors.Is(err, net.ErrClosed) &&
			!errors.Is(err, http.ErrServerClosed) {
			slog.Error("HTTP server error", "error", err)
		}
	}()

	slog.Info("listening on NetherNet", "address", tcp.Addr())

	conn, err := l.Accept()
	if err != nil {
		_ = srv.Close()
		_ = tcp.Close()
		_ = l.Close()
		return nil, err
	}

	slog.Info("connected",
		"remoteAddr", conn.RemoteAddr(),
		"localAddr", conn.LocalAddr(),
		"latency", conn.(*nethernet.Conn).Latency(),
	)
	return conn.(*nethernet.Conn), nil

}
