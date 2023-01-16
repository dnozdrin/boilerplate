package order

import (
	"github.com/dnozdrin/boilerplate/domain"
)

type Event interface {
	domain.Event
	OrderID() string
}

type CreatedEvent struct {
	orderID    string
	customerID string
	lineItems  []LineItem
}

func (e CreatedEvent) Name() string {
	return "event.order.created"
}

func (e CreatedEvent) IsAsynchronous() bool {
	return false
}

func (e CreatedEvent) OrderID() string {
	return e.orderID
}

func (e CreatedEvent) CustomerID() string {
	return e.customerID
}

func (e CreatedEvent) LineItems() []LineItem {
	return e.lineItems
}

type AcceptedEvent struct {
	orderID   string
	managerID string
}

func (e AcceptedEvent) Name() string {
	return "event.order.accepted"
}

func (e AcceptedEvent) IsAsynchronous() bool {
	return false
}

func (e AcceptedEvent) OrderID() string {
	return e.orderID
}

func (e AcceptedEvent) ManagerID() string {
	return e.managerID
}

type CompletedEvent struct {
	orderID string
}

func (e CompletedEvent) Name() string {
	return "event.order.completed"
}

func (e CompletedEvent) IsAsynchronous() bool {
	return false
}

func (e CompletedEvent) OrderID() string {
	return e.orderID
}

type ReassignedEvent struct {
	orderID       string
	managerID     string
	prevManagerID string
}

func (e ReassignedEvent) Name() string {
	return "event.order.reassigned"
}

func (e ReassignedEvent) IsAsynchronous() bool {
	return false
}

func (e ReassignedEvent) OrderID() string {
	return e.orderID
}

func (e ReassignedEvent) ManagerID() string {
	return e.managerID
}

func (e ReassignedEvent) PrevManagerID() string {
	return e.prevManagerID
}

type LineItemsAddedEvent struct {
	orderID   string
	lineItems []LineItem
}

func (e LineItemsAddedEvent) Name() string {
	return "event.order.itemsAdded"
}

func (e LineItemsAddedEvent) IsAsynchronous() bool {
	return false
}

func (e LineItemsAddedEvent) OrderID() string {
	return e.orderID
}

func (e LineItemsAddedEvent) LineItems() []LineItem {
	return e.lineItems
}

type CancelledEvent struct {
	orderID   string
	lineItems []LineItem
}

func (e CancelledEvent) Name() string {
	return "event.order.cancelled"
}

func (e CancelledEvent) IsAsynchronous() bool {
	return false
}

func (e CancelledEvent) OrderID() string {
	return e.orderID
}
