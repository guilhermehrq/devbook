package models

import "time"

// User is a struct for a user in the database
type User struct {
	ID        uint       `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Nickname  string     `json:"nickname,omitempty"`
	Email     string     `json:"email,omitempty"`
	Password  string     `json:"password,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}
