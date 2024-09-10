// models/models.go
package models

type User struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"-"`
}