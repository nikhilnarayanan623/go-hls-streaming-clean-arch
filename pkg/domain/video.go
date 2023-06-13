package domain

import "time"

type Video struct {
	ID          string `json:"id" gorm:"primaryKey;not null"`
	Name        string `json:"name" gorm:"not null"`
	VideoUrl    string `json:"video_url" gorm:"not null"`
	PlaylistUrl string `json:"playlist_url" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	//Duration    time.Duration `json:"duration" gorm:"not null"`
	UploadedAt time.Time `json:"uploaded_at" gorm:"not null"`
}
