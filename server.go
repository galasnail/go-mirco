package gomicro

import (
	"time"
	"github.com/hashicorp/net-rpc-msgpackrpc"
	"google.golang.org/grpc"
	"net"
	"io"
)


// ServerCodecFunc is used to create a rpc.ServerCodec from net.Conn.
type ServerCodecFunc func(conn io.ReadWriteCloser) grpc.Codec

type Server struct {
	ServerCodecFunc ServerCodecFunc
	//PluginContainer must be configured before starting and Register plugins must be configured before invoking RegisterName method
	PluginContainer IServerPluginContainer
	//Metadata describes extra info about this service, for example, weight, active status
	Metadata     string
	rpcServer    *grpc.Server
	listener     net.Listener
	Timeout      time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// NewServer returns a new Server.
func NewServer() *Server {
	return &Server{
		rpcServer:       grpc.NewServer(),
		PluginContainer: &ServerPluginContainer{make([]IPlugin, 0)},
		ServerCodecFunc: msgpackrpc.NewServerCodec,
	}
}
