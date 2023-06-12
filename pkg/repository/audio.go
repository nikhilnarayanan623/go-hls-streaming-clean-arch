package repository

import (
	"github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type audioDatabase struct {
	db *gorm.DB
}

func NewAudioRepository(db *gorm.DB) interfaces.AudioRepository {
	return &audioDatabase{
		db: db,
	}
}
