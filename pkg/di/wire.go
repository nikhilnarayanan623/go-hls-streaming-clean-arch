//go:build wireinject

package di

import (
	"github.com/google/wire"
	http "github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/api"
	"github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/api/handler"
	"github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/config"
	"github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/db"
	"github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/repository"
	"github.com/nikhilnarayanan623/go-hls-streaming-clean-arch/pkg/usecase"
)

func InitializeApi(cfg config.Config) (*http.Server, error) {

	wire.Build(
		db.ConnectDatabase,
		repository.NewVideoRepository,
		repository.NewAudioRepository,
		usecase.NewVideoUseCase,
		usecase.NewAudioUseCase,
		handler.NewVideoHandler,
		handler.NewAudioHandler,
		http.NewServerHTTP,
	)

	return &http.Server{}, nil
}
