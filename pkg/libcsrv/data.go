package libcsrv

// Path to server socket
const SocketPath = "/srv/rootd"

// Rootexec data
const (
	RootexecPath         = "/usr/local/bin/rootexec"
	RootexecFlagName     = "-n"
	RootexecFlagOverride = "-c"
)

// WebAPI Routes
const (
	RouteStart      = "/api/start"
	RouteStop       = "/api/stop"
	RouteListAll    = "/api/list/all"
	RouteListOnline = "/api/list/online"
	RoutePause      = "/api/pause"
	RouteResume     = "/api/resume"
	RouteOnline     = "/api/online"
)
