package order

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	id         string
	customerID string
	managerID  string
	status     Status
	lineItems  []LineItem
	createdAt  time.Time
	updatedAt  time.Time
}

func NewOrder(customerID string, items []LineItem) (Order, CreatedEvent) {
	order := Order{
		id:         uuid.New().String(),
		customerID: customerID,
		status:     New,
		lineItems:  items,
		createdAt:  time.Now().UTC(),
		updatedAt:  time.Now().UTC(),
	}

	return order, CreatedEvent{
		orderID:    order.id,
		customerID: order.customerID,
		lineItems:  order.lineItems, // todo: make sure it's really needed
	}
}

func (o *Order) ID() string {
	return o.id
}

func (o *Order) CustomerID() string {
	return o.customerID
}

func (o *Order) ManagerID() string {
	return o.managerID
}

func (o *Order) Status() Status {
	return o.status
}

func (o *Order) LineItems() []LineItem {
	return o.lineItems
}

func (o *Order) CreatedAt() time.Time {
	return o.createdAt
}

func (o *Order) UpdatedAt() time.Time {
	return o.updatedAt
}

func (o *Order) Accept(managerID string) AcceptedEvent {
	o.status = Accepted
	o.managerID = managerID
	o.updatedAt = time.Now().UTC()

	return AcceptedEvent{
		orderID:   o.id,
		managerID: o.managerID,
	}
}

// todo: maybe interface and error = event
func (o *Order) AddLineItems(items ...LineItem) (LineItemsAddedEvent, error) {
	if o.status != New {
		return LineItemsAddedEvent{}, ErrInvalidStatus
	}

	o.lineItems = append(o.lineItems, items...)
	o.updatedAt = time.Now().UTC()

	return LineItemsAddedEvent{
		orderID:   o.id,
		lineItems: o.lineItems,
	}, nil
}

func (o *Order) Reassign(managerID string) ReassignedEvent {
	event := ReassignedEvent{
		orderID:       o.id,
		managerID:     managerID,
		prevManagerID: o.managerID,
	}

	o.managerID = managerID
	o.updatedAt = time.Now().UTC()

	return event
}

func (o *Order) Complete() CompletedEvent {
	o.status = Completed
	o.updatedAt = time.Now().UTC()

	return CompletedEvent{orderID: o.id}
}

func (o *Order) Cancel() CancelledEvent {
	o.status = Cancelled
	o.updatedAt = time.Now().UTC()

	return CancelledEvent{orderID: o.id}
}

type LineItem struct {
	ID        string
	ProductID string
	Quantity  int
}
