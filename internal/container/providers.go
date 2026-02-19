package container

import (
	"main/internal/adapter/protocol"
	"main/internal/adapter/protocol/handler"
	"main/internal/domain/repository"
	"main/internal/usecase"
)

type Container struct {
	Store          repository.KeyValueRepository
	Persistence    repository.PersistenceRepository
	CommandHandler *usecase.CommandHandler
	TCPHandler     *handler.TCPHandler
	Parser         *protocol.Parser
}

func NewContainer(
	store repository.KeyValueRepository,
	persistence repository.PersistenceRepository,
	commandHandler *usecase.CommandHandler,
	tcpHandler *handler.TCPHandler,
	parser *protocol.Parser,
) *Container {
	return &Container{
		Store:          store,
		Persistence:    persistence,
		CommandHandler: commandHandler,
		TCPHandler:     tcpHandler,
		Parser:         parser,
	}
}

func (c *Container) Close() error {
	if c.Persistence != nil {
		return c.Persistence.Close()
	}

	return nil
}
