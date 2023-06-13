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

// Upload godoc
// @summary api for upload videos to server
// @tags Video
// @id Upload
// @Param     video   formData     file   true   "Video file to upload"
// @Param     name   formData     string   true   "Video Name"
// @Param     description   formData     string   true   "Video Description"
// @Router /video [post]
// @Success 201 {object} response.Response{} "successfully video saved"
// @Failure 400 {object} response.Response{}  "failed get inputs"
// @Failure 500 {object} response.Response{}  "failed to save video"
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
		response.ErrorResponse(ctx, http.StatusBadRequest, "failed to find form values from request", err, nil)
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

// FindAll godoc
// @summary api for find all videos on server
// @tags Video
// @id FindAll
// @Param     page_number   query     string   false   "Page Number"
// @Param     count   query     string   false   "Count"
// @Router /video/all [get]
// @Success 201 {object} response.Response{} "successfully found all videos"
// @Failure 500 {object} response.Response{}  "failed to get all videos"
func (c *videoHandler) FindAll(ctx *gin.Context) {

	pagination := request.GetPagination(ctx)

	videos, err := c.videoUseCase.FindAll(ctx, pagination)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "failed to get all videos", err, nil)
		return
	}

	if videos == nil {
		response.SuccessResponse(ctx, http.StatusOK, "there is no videos to show")
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, "successfully found all videos", videos)
}
