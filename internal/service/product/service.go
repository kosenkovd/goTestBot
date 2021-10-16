package product

var allProducts = []Product{
	{
		Title: "first",
	},
	{
		Title: "second",
	},
	{
		Title: "third",
	},
	{
		Title: "fourth",
	},
	{
		Title: "fifth",
	},
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProducts
}
