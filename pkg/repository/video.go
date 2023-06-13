package repository

import (
	"context"
	"time"

	"github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/domain"
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

func (c *videoDatabase) Save(ctx context.Context, video domain.Video) error {

	uploadedAt := time.Now()
	query := `INSERT INTO videos 
	(id, name, video_url, playlist_url, description, uploaded_at) 
	VALUES ($1, $2, $3, $4, $5, $6)`
	err := c.db.Exec(query, video.ID, video.Name, video.VideoUrl,
		video.PlaylistUrl, video.Description, uploadedAt).Error

	return err
}

func (c *videoDatabase) FindByID(ctx context.Context, id string) (video domain.Video, err error) {

	query := `SELECT id, name, video_url, playlist_url, description, uploaded_at 
	FROM videos WHERE id = $1`
	err = c.db.Raw(query, id).Scan(&video).Error

	return
}
