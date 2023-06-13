package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/domain"
)

type VideoRepository interface {
	Save(ctx context.Context, video domain.Video) error
	FindByID(ctx context.Context, id string) (video domain.Video, err error)
}
