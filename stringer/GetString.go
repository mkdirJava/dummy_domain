package stringer

import "fmt"

type GetStringer interface {
	GetString() string
	Greeting(string) string
}

func NewGetStringerService() GetStringer {
	return &GetStringerService{}
}

type GetStringerService struct{}

func (g *GetStringerService) GetString() string {
	return "hi"
}

func (g *GetStringerService) Greeting(arg string) string {
	return fmt.Sprintf("Hi there %s", arg)
}
