package order

import (
	"context"
	domain "github.com/dnozdrin/boilerplate/domain/order"
	"github.com/dnozdrin/boilerplate/outbound/postgres"
)

type Repo struct {
	db postgres.DB
}

func NewRepo(db postgres.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) FindByID(ctx context.Context, id string) (domain.Order, error) {
	var order domain.Order
	err := r.db.WithConn(ctx).
		NewSelect().
		Model(order).
		Where("id = ?", id).
		Scan(ctx)

	return order, err
}

func (r *Repo) Save(ctx context.Context, order domain.Order) error {
	_, err := r.db.WithConn(ctx).
		NewInsert().
		Model(order).
		On("CONFLICT (id) DO UPDATE").
		Exec(ctx)

	return err
}
