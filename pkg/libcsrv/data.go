package libcsrv

// Path to server socket
const SocketPath = "/net/rootd"

// Rootexec data
const (
	RootexecPath = "/usr/local/bin/rootexec"
	RootexecArg  = "-n"
)

// States
const (
	StateOff    = "Offline"
	StateOn     = "Online"
	StatePaused = "Paused"
)

// WebAPI Routes
const (
	RouteStart      = "/api/start"
	RouteStop       = "/api/stop"
	RouteListAll    = "/api/list/all"
	RouteListOnline = "/api/list/online"
	RoutePause      = "/api/pause"
	RouteResume     = "/api/resume"
)
