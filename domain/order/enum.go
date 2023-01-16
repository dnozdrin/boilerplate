package order

type Status int

const (
	_ Status = iota
	New
	Accepted
	Cancelled
	Completed
)

var statuses = [...]string{
	"UNKNOWN",
	"NEW",
	"ACCEPTED",
	"CANCELLED",
	"COMPLETED",
}

// todo: fix write to DB

func (s Status) String() string {
	return statuses[s]
}
