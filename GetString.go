package getstring

type GetStringer interface {
	GetString() string
}

func NewGetStringerService() GetStringer {
	return &GetStringerService{}
}

type GetStringerService struct{}

func (g *GetStringerService) GetString() string {
	return "hi"
}
