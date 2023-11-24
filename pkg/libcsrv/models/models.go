package models

type VeEntry struct {
	Name    string
	State   string
	Command string
}

type VeList []VeEntry
