package libcsrv

import (
	"net"
	"net/http"
)

// Allocated and initializes http client over unix socket
func MakeClient() http.Client {
	return http.Client{
		Transport: &http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				return net.Dial("unix", SocketPath)
			},
		},
	}
}

// Append the protocol prefix
func MapProtocol(s string) string {
	return "http://localhost" + s
}
