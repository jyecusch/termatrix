package matrix

type RuneStatus int

const (
	RuneStatusInactive RuneStatus = iota
	RuneStatusHead
	RuneStatusActive
	RuneStatusTail
)
