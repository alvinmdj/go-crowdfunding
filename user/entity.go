package user

import "time"

type User struct {
	ID             int       `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name"`
	Occupation     string    `json:"occupation"`
	Email          string    `json:"email"`
	PasswordHash   string    `json:"-"`
	AvatarFileName string    `json:"avatar_file_name"`
	Role           string    `json:"role"`
	Token          string    `json:"token"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
