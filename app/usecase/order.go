package usecase

import (
	"context"

	"github.com/dnozdrin/boilerplate/app"
	domain "github.com/dnozdrin/boilerplate/domain/order"
)

type Service struct {
	repository Repository
	transactor app.Transactor
	publisher  app.EventPublisher
}

func NewService(
	repository Repository,
	transactor app.Transactor,
	publisher app.EventPublisher,
) *Service {
	return &Service{
		repository: repository,
		transactor: transactor,
		publisher:  publisher,
	}
}

type Repository interface {
	Save(ctx context.Context, order domain.Order) error
	FindByID(ctx context.Context, id string) (domain.Order, error)
}

func (s *Service) Create(ctx context.Context, customerID string, items []domain.LineItem) (domain.Order, error) {
	order, event := domain.NewOrder(customerID, items)
	if err := s.transactor.WithinTransaction(ctx, func(ctx context.Context) error {
		if err := s.repository.Save(ctx, order); err != nil {
			return err
		}

		s.publisher.Notify(ctx, event)

		return nil
	}); err != nil {
		return domain.Order{}, err
	}

	return order, nil
}

func (s *Service) Accept(ctx context.Context, orderID, managerID string) error {
	order, err := s.repository.FindByID(ctx, orderID)
	if err != nil {
		return err
	}

	event := order.Accept(managerID)
	if err = s.repository.Save(ctx, order); err != nil {
		return err
	}

	s.publisher.Notify(ctx, event)

	return nil
}

func (s *Service) AddLineItems(ctx context.Context, orderID string, items []domain.LineItem) error {
	order, err := s.repository.FindByID(ctx, orderID)
	if err != nil {
		return err
	}

	event, err := order.AddLineItems(items...)
	if err != nil {
		return err
	}

	if err = s.repository.Save(ctx, order); err != nil {
		return err
	}

	s.publisher.Notify(ctx, event)

	return nil
}

func (s *Service) Reassign(ctx context.Context, orderID, managerID string) error {
	order, err := s.repository.FindByID(ctx, orderID)
	if err != nil {
		return err
	}

	event := order.Reassign(managerID)
	if err = s.repository.Save(ctx, order); err != nil {
		return err
	}

	s.publisher.Notify(ctx, event)

	return nil
}

func (s *Service) Complete(ctx context.Context, orderID string) error {
	order, err := s.repository.FindByID(ctx, orderID)
	if err != nil {
		return err
	}

	event := order.Complete()
	if err = s.repository.Save(ctx, order); err != nil {
		return err
	}

	s.publisher.Notify(ctx, event)

	return nil
}

func (s *Service) Cancel(ctx context.Context, orderID string) error {
	order, err := s.repository.FindByID(ctx, orderID)
	if err != nil {
		return err
	}

	event := order.Cancel()
	if err = s.repository.Save(ctx, order); err != nil {
		return err
	}

	s.publisher.Notify(ctx, event)

	return nil
}
