package models

import (
	"time"

	"me.tofaa/internal/pkg/optional"
)

type Permission byte

const (
	PermissionAdmin Permission = 1 << iota
)

type User struct {
	ID          uint64                       `json:"id"`
	Username    string                       `json:"username"`
	Password    string                       `json:"password"`
	Email       string                       `json:"email"`
	CreatedAt   time.Time                    `json:"created_at"`
	UpdatedAt   optional.Optional[time.Time] `json:"updated_at"`
	Permissions Permission
}

func (u *User) IsAdmin() bool {
	return u.Permissions&PermissionAdmin != 0
}
