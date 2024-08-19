package interfaces

type Data interface {
	GetKey() string
	GetSubData() []*Data
	SetSubData([]*Data) error
	GetData() string
	SetData(string) error
}

type Database interface {
	GetData(string) (Data, error)
	SetData(string, Data) error
}
