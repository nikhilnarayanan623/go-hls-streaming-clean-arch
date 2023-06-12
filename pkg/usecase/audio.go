package usecase

import (
	repo "github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/repository/interfaces"
	"github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/usecase/interfaces"
)

type audioUseCase struct {
	audioRepo repo.AudioRepository
}

func NewAudioUseCase(audioRepo repo.AudioRepository) interfaces.AudioUseCase {
	return &audioUseCase{
		audioRepo: audioRepo,
	}
}
