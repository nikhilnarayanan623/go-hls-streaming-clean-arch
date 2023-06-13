package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/utils/request"
	"github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/utils/response"
)

type VideoUseCase interface {
	Save(ctx context.Context, req request.UploadVideo) (videoID string, err error)
	FindAll(ctx context.Context, pagination request.Pagination) (videos []response.VideoDetails, err error)
}
