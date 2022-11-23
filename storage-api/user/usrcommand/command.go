package usrcommand

type Command interface {
	Execute(data []byte) error
}
