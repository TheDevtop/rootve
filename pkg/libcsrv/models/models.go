package models

type VeEntry struct {
	Name    string
	Path    string
	State   string
	Command string
}

type VeList []VeEntry
