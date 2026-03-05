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

func (s Status) IsValid() bool {
	switch s {
	case StatusPending, StatusActive, StatusApproved, StatusRejected, StatusClosed, StatusInactive:
		return true
	}
	return false
}
func (s Status) String() string {
	return string(s)
}
