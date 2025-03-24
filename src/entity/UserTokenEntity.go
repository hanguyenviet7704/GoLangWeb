package entity

import "time"

type UserToken struct {
	ID           int    `json:"id"`
	UserID       int    `json:"user_id"`
	DeviceID     string `json:"device_id"`
	AccessToken  string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    time.Time
	CreatedAt    time.Time
}
