package product

import "fmt"

type Product struct {
	Title string
}

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

func (p *Product) String() string {
	return fmt.Sprintf("\"%v\"", p.Title)
}
