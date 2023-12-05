package libcsrv

type State struct {
	Byte  byte
	Label string
}

// States
const (
	StateOff    byte = 0x00
	StateOn     byte = 0xff
	StatePaused byte = 0x0f
)

// State labels
const (
	SlabelOff    = "Offline"
	SlabelOn     = "Online"
	SlabelPaused = "Paused"
)

// Convert state to label
func Slabel(state byte) string {
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
