package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/utils/request"
)

type VideoUseCase interface {
	Save(ctx context.Context, req request.UploadVideo) (videoID string, err error)
}
