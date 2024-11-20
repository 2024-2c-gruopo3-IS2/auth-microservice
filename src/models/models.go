// models/models.go
package models

type User struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"-"`
	IsBlocked bool `db:"is_blocked" json:"-"`
}

type UserResponse struct {
	Email    string `db:"email" json:"email"`
	IsBlocked bool `db:"is_blocked" json:"is_blocked"`
}

type ResetResponse struct {
	Token string `db:"token" json:"token"`
	CreatedAt string `db:"created_at" json:"created_at"`
}

type PinResponse struct {
	Pin string `db:"pin" json:"pin"`
	CreatedAt string `db:"created_at" json:"created_at"`
}

type Block struct {
	Reason string `db:"reason" json:"reason"`
	Days int `db:"days" json:"days"`
	CreatedAt string `db:"created_at" json:"created_at"`
}