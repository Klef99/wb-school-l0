//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
)

func InitializeDependencies() (*Command, func(), error) {
	panic(
		wire.Build(
			InfrastructureSet,
			HHTPAdapterSet,
			StorageSet,
			ServiceSet,
			KafkaAdapterSet,
			HTTPHandlerSet,
			ProvideCommand,
		),
	)
}
