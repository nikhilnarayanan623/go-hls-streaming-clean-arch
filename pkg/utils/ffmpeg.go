package utils

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// create new playlist for file using ffmpeg on given playlistDir
func CreatePlaylistUsingFfmpeg(playlistDir, inputFilePath string, segmentDuration time.Duration) (err error) {

	if err := os.MkdirAll(playlistDir, 0700); err != nil {
		return fmt.Errorf("failed to create playlist directory \nerror:%w", err)
	}

	// playlistPath = fmt.Sprintf("%s/playlist.m3u8", outputDir)

	exec := exec.Command("ffmpeg",
		"-i", inputFilePath,
		"-profile:v", "baseline",
		"-level", "3.0",
		"-start_number", "0",
		"-hls_time", segmentDuration.String(), // duration of the segment
		"-hls_list_size", "0",
		"-f", "hls",
		fmt.Sprintf("%s/playlist.m3u8", playlistDir),
	)

	_, err = exec.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to run ffmpeg command for create playlist \nerror:%w", err)
	}

	return nil
}
