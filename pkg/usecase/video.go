package usecase

import (
	repo "github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/repository/interfaces"
	"github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/usecase/interfaces"
)

type videoUseCase struct {
	videoRepo repo.VideoRepository
}

func NewVideoUseCase(videoRepo repo.VideoRepository) interfaces.VideoUseCase {
	return &videoUseCase{
		videoRepo: videoRepo,
	}
}
