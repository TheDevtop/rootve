package libcsrv

// Path to server socket
const SocketPath = "/net/rootd"

// States
const (
	StateOff = "Offline"
	StateOn  = "Online"
)

// WebAPI Routes
const (
	RouteStart      = "/api/start"
	RouteStop       = "/api/stop"
	RouteListAll    = "/api/list/all"
	RouteListOnline = "/api/list/online"
)
