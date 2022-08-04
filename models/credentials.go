package models

import "time"

type Credential struct {
	Id        int
	Role      string
	Email     string
	Password  string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
