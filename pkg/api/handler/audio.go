package handler

import (
	"github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/api/handler/interfaces"
	usecase "github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/usecase/interfaces"
)

type audioHandler struct {
	audioUseCase usecase.AudioUseCase
}

func NewAudioHandler(audioUseCase usecase.AudioUseCase) interfaces.AudioHandler {
	return &audioHandler{
		audioUseCase: audioUseCase,
	}
}
