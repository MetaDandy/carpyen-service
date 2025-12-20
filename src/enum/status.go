package enum

type Status string

const (
	StatusPending  Status = "PENDING"
	StatusActive   Status = "ACTIVE"
	StatusApproved Status = "APPROVED"
	StatusRejected Status = "REJECTED"
	StatusClosed   Status = "CLOSED"
	StatusInactive Status = "INACTIVE"
)
