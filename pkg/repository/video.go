package repository

import (
	"github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type videoDatabase struct {
	db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) interfaces.VideoRepository {
	return &videoDatabase{
		db: db,
	}
}
