package handler

import(
	"github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/api/handler/interfaces"
	usecase "github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/usecase/interfaces"
) 

type videoHandler struct {
	videoUseCase usecase.VideoUseCase
}

func NewVideoHandler(videoUseCase usecase.VideoUseCase) interfaces.VideoHandler {
	return &videoHandler{
		videoUseCase: videoUseCase,
	}
}
