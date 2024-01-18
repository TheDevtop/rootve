package libcsrv

// Generic base form
type Form[T any] struct {
	Error bool
	Data  T
}

// Virtual Environment list form
type FormVeList struct {
	Name      string
	State     string
	Interface string
	Address   string
	Command   string
}

// Basic message form
type FormMessage Form[string]
