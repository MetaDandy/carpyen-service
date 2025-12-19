package enum

type StatusEnum string

const (
	StatusPending  StatusEnum = "PENDING"
	StatusActive   StatusEnum = "ACTIVE"
	StatusApproved StatusEnum = "APPROVED"
	StatusRejected StatusEnum = "REJECTED"
	StatusClosed   StatusEnum = "CLOSED"
	StatusInactive StatusEnum = "INACTIVE"
)
