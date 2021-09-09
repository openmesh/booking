package booking

// Block represents a period of time that a resource has been marked as
// unavailable for. This can be used to indicate that a resource should
// temporarily stop operating according to its normal defined slots. Useful for
// holidays etc.
type Block struct {
	ID int `json:"id"`
}
