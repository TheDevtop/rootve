package libcsrv

import "errors"

// Path to server socket
const SocketPath = "/net/rootd"

// Rootexec data
const (
	RootexecPath         = "/usr/local/bin/rootexec"
	RootexecFlagName     = "-n"
	RootexecFlagOverride = "-c"
)

// Request header keys
const HdrName = "name"

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

// Internal errors
var errInvalidState = errors.New("invalid target state")
