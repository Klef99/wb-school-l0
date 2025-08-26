//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
)

func InitializeDependencies() (HTTPAdapter, func(), error) {
	panic(
		wire.Build(
			InfrastructureSet,
			StorageSet,
			ServiceSet,
			HHTPAdapterSet,
			HTTPHandlerSet,
		),
	)
}
