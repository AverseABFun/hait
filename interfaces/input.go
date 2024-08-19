package interfaces

type Inputter interface {
	HasInput() bool
	GetInput() (string, error)
	DisableInput() error
	EnableInput() error
}
