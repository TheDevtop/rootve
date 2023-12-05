package libcsrv

type Form[T any] struct {
	Error bool
	Data  T
}

type FormVeList struct {
	Name    string
	State   string
	Path    string
	Command string
}

type FormMessage Form[string]
