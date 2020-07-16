package models

type User struct {
	ID                     int
	Name                   string
	Email                  string
	Photo                  string
	Password               string
	Password_changed_at    float64
	Password_reset_token   string
	Password_reset_expires float64
}
