package services

type Service struct {
	Logger   Logger
	NService NService
}

func ProvideService(l Logger, n NService) Service {
	return Service{
		Logger:   l,
		NService: n,
	}
}

func (s *Service) GetData() string {
	s.Logger.Logged()

	return s.NService.NMethod("test")
}
