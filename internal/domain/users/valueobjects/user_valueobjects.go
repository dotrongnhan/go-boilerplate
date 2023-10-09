package valueobjects

import "time"

// UserFilter represents a filter for querying users.
type UserFilter struct {
	ID        *int64
	Name      *string
	Email     *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

// UserSort represents a sort order for querying users.
type UserSort struct {
	ID        SortDirection
	Name      SortDirection
	Email     SortDirection
	CreatedAt SortDirection
	UpdatedAt SortDirection
}

// SortDirection represents a sort direction.
type SortDirection int

const (
	// Ascending represents an ascending sort direction.
	Ascending SortDirection = iota
	// Descending represents a descending sort direction.
	Descending
)
