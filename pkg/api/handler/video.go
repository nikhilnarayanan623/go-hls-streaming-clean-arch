package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/api/handler/interfaces"
	usecase "github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/usecase/interfaces"
	"github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/utils/request"
	"github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/utils/response"
)

type videoHandler struct {
	videoUseCase usecase.VideoUseCase
}

func NewVideoHandler(videoUseCase usecase.VideoUseCase) interfaces.VideoHandler {
	return &videoHandler{
		videoUseCase: videoUseCase,
	}
}

func (c *videoHandler) Upload(ctx *gin.Context) {

	fileHeader, err := ctx.FormFile("video")
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, "failed get video from request", err, nil)
		return
	}
	name, err1 := request.GetFormValues(ctx, "name")
	description, err2 := request.GetFormValues(ctx, "description")
	err = errors.Join(err1, err2)

	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, "failed to form values from request", err, nil)
		return
	}

	videoID, err := c.videoUseCase.Save(ctx, request.UploadVideo{
		Name:        name,
		Description: description,
		FileHeader:  fileHeader,
	})

	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "failed to save video", err, nil)
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, "successfully video saved", gin.H{
		"video_id": videoID,
	})
}
