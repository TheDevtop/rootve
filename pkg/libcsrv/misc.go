package libcsrv

// Path to server socket
const SocketPath = "/srv/rootd"

// Web API routes
const (
	RouteStart      = "/api/start"
	RouteStop       = "/api/stop"
	RouteListAll    = "/api/list/all"
	RouteListOnline = "/api/list/online"
	RoutePause      = "/api/pause"
	RouteResume     = "/api/resume"
	RouteRemove     = "/api/remove"
)
