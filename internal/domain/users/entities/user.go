package entities

import "time"

// UserStatus represents the status of a user.
type UserStatus string

const (
	// UserStatusActive represents an active user.
	UserStatusActive UserStatus = "active"
	// UserStatusInactive represents an inactive user.
	UserStatusInactive UserStatus = "inactive"
)

// User represents a user in the system.
type User struct {
	ID         int64      `json:"id"`
	Username   string     `json:"username"`
	Password   string     `json:"password"`
	Email      string     `json:"email"`
	Status     UserStatus `json:"status"`
	CreateTime time.Time  `json:"createTime"`
	UpdateTime time.Time  `json:"updateTime"`
}
