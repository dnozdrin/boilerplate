package order

import (
	"github.com/dnozdrin/errdetail"
)

var (
	ErrInvalidStatus = errdetail.NewFailedPrecondition("invalid order status") // todo: review error
)
