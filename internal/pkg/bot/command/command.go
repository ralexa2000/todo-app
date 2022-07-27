package command

type Interface interface {
	Name() string
	Arguments() string
	Description() string
	Process(userName string, inputString string) string
}
