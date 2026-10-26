package gateway

type server struct {
	name    string
	address string
}

func (proxy *proxy) RegisterDefaultServer(name, addr string) *server {
	return &server{}
}

func (proxy *proxy) RegisterServer(name, addr string) {

}

func getServer() (addr string) {
	return ":19133"
}
