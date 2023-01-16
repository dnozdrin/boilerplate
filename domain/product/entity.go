package product

import (
	"github.com/dnozdrin/boilerplate/domain"
)

type Product struct {
	id    string
	title string
	price domain.Money
}
