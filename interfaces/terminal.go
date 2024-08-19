package interfaces

type Command interface {
	GetCallName() string
	GetBriefDescription() string
	GetDescription() string
}

type Terminal interface {
	SendInput(string) error
	GetNewOutput() (string, error)
	HasNewOutput() (bool, error)
	GetCommands() []Command
	SendFeedback(uint8) error
}
