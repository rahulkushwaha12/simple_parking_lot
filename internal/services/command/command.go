package command

type ICommandService interface {
	Parse(string) error
	Run() string
}
