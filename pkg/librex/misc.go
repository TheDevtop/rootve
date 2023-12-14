package librex

// States
const (
	StateOff    byte = 0x00
	StateOn     byte = 0xff
	StatePaused byte = 0x0f
)

// Rootexec data
const (
	RootexecPath         = "/usr/local/bin/rootexec"
	RootexecFlagName     = "-n"
	RootexecFlagOverride = "-c"
)

// State labels
const (
	SlabelOff    = "Offline"
	SlabelOn     = "Online"
	SlabelPaused = "Paused"
)

// Convert state to label
func StateToLabel(state byte) string {
	switch state {
	case StateOff:
		return SlabelOff
	case StateOn:
		return SlabelOn
	case StatePaused:
		return SlabelPaused
	}
	return ""
}
