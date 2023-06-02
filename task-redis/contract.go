package user_queue

import "context"

// struct for User
type User struct {
	ID       string
	Username string
	Password string
}

// initialization interface
type UserQueue interface {
	Count(ctx context.Context) int64
	PushQueue(ctx context.Context, user *User) error
	PopQueue(ctx context.Context) (*User, error)
}