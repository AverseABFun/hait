package interfaces

type MessageId string

type Outputter interface {
	OutputText(string) (error, MessageId)
	EditText(MessageId, string) error
}
