package entity

import "time"

type SocialMedia struct {
	ID             uint      `json:"id gorm:primaryKey"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Name           string    `json:"name" gorm:"unique;not null"`
	SocialMediaUrl string    `json:"social_media_url" gorm:"unique;not null"`
	UserID         uint      `json:"user_id"`
	User           User      `json:"user"`
}
