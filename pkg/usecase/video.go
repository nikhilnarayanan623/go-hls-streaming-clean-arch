package usecase

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/domain"
	repo "github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/repository/interfaces"
	"github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/usecase/interfaces"
	"github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/utils"
	"github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/utils/request"
)

type videoUseCase struct {
	videoRepo repo.VideoRepository
}

const (
	uploadDir       = "./storage/uploads"
	playListDir     = "./storage/playlist"
	videoExtension  = ".mp4"
	segmentDuration = 10 * time.Second
)

func NewVideoUseCase(videoRepo repo.VideoRepository) interfaces.VideoUseCase {
	return &videoUseCase{
		videoRepo: videoRepo,
	}
}

func (c *videoUseCase) Save(ctx context.Context, req request.UploadVideo) (string, error) {

	videoID := utils.GenerateUniqueString()
	// create a new video file for copy the request file
	file, videoFullPath, err := c.createANewVideoFile(videoID)
	if err != nil {
		return "", fmt.Errorf("failed to create a new video file \nerror:%w", err)
	}

	reqFile, err := req.FileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open request file \nerror:%w", err)
	}
	// copy the request file into created video file
	if err = utils.CopyFilesFormSrcToDest(reqFile, file); err != nil {
		return "", fmt.Errorf("failed copy file from request \nerror:%s", err)
	}

	// create playlist
	outputDir := playListDir + "/" + videoID
	playListPath, err := utils.CreatePlaylistUsingFfmpeg(outputDir, videoFullPath, segmentDuration)
	if err != nil {
		return "", fmt.Errorf("failed to create playlist \nerror:%w", err)
	}

	// save video details on database
	err = c.videoRepo.Save(ctx, domain.Video{
		ID:          videoID,
		Name:        req.Name,
		Description: req.Description,
		VideoUrl:    videoFullPath,
		PlaylistUrl: playListPath,
	})

	if err != nil {
		return "", fmt.Errorf("failed to save video details on database")
	}

	return videoID, nil
}

// create new video file using video id
func (c *videoUseCase) createANewVideoFile(videoID string) (file *os.File, fileFullPath string, err error) {

	fileDir := uploadDir + "/" + videoID + "/"
	fileFullPath = fileDir + videoID + videoExtension

	if err = os.MkdirAll(fileDir, 0700); err != nil {
		return nil, "", fmt.Errorf("failed to create file directory \nerror:%w", err)
	}

	if file, err = os.Create(fileFullPath); err != nil {
		return nil, "", fmt.Errorf("failed to create a new video file \nerror:%w", err)
	}

	return
}
