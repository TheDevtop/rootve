package libcsrv

type Form[T any] struct {
	Error   bool
	Message string
	Data    T
}
