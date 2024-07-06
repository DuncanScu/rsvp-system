package status

type StatusType string

const (
	Accepted StatusType = "Accepted"
	Declined StatusType = "Declined"
	Pending  StatusType = "Pending"
)
