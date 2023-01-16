package domain

type Event interface {
	Name() string
	IsAsynchronous() bool
}
