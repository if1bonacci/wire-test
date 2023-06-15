package services

type NService struct{}

func ProvideNService() NService {
	return NService{}
}

func (n *NService) NMethod(s string) string {
	return s
}
