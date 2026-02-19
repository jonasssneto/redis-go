//go:build wireinject
// +build wireinject

package container

import (
	"main/internal/adapter/protocol"
	"main/internal/adapter/protocol/handler"
	"main/internal/infraestructure/persistence"
	"main/internal/infraestructure/storage"
	"main/internal/usecase"

	"github.com/google/wire"
)

func InitializeContainer(opt persistence.AOFProviderOption) (*Container, func(), error) {
	wire.Build(
		// Infraestructure providers
		storage.NewStore,
		persistence.NewAOFProvider,

		// Adapter providers
		protocol.NewParser,

		// Use case providers
		usecase.NewStats,
		usecase.NewCommandHandler,

		// Handler providers
		handler.NewTCPHandler,

		// Container provider
		NewContainer,
	)

	return nil, nil, nil
}
