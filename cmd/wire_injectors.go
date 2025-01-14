//+build wireinject

package cmd

import (
	"sync"

	"github.com/google/wire"
	"github.com/navidrome/navidrome/core"
	"github.com/navidrome/navidrome/persistence"
	"github.com/navidrome/navidrome/scanner"
	"github.com/navidrome/navidrome/scheduler"
	"github.com/navidrome/navidrome/server"
	"github.com/navidrome/navidrome/server/events"
	"github.com/navidrome/navidrome/server/nativeapi"
	"github.com/navidrome/navidrome/server/subsonic"
)

var allProviders = wire.NewSet(
	core.Set,
	subsonic.New,
	nativeapi.New,
	persistence.New,
	GetBroker,
)

func CreateServer(musicFolder string) *server.Server {
	panic(wire.Build(
		server.New,
		allProviders,
	))
}

func CreateNativeAPIRouter() *nativeapi.Router {
	panic(wire.Build(
		allProviders,
	))
}

func CreateSubsonicAPIRouter() *subsonic.Router {
	panic(wire.Build(
		allProviders,
		GetScanner,
	))
}

// Scanner must be a Singleton
var (
	onceScanner     sync.Once
	scannerInstance scanner.Scanner
)

func GetScanner() scanner.Scanner {
	onceScanner.Do(func() {
		scannerInstance = createScanner()
	})
	return scannerInstance
}

func createScanner() scanner.Scanner {
	panic(wire.Build(
		allProviders,
		scanner.New,
	))
}

// Broker must be a Singleton
var (
	onceBroker     sync.Once
	brokerInstance events.Broker
)

func GetBroker() events.Broker {
	onceBroker.Do(func() {
		brokerInstance = createBroker()
	})
	return brokerInstance
}

func createBroker() events.Broker {
	panic(wire.Build(
		events.NewBroker,
	))
}

// Scheduler must be a Singleton
var (
	onceScheduler     sync.Once
	schedulerInstance scheduler.Scheduler
)

func GetScheduler() scheduler.Scheduler {
	onceScheduler.Do(func() {
		schedulerInstance = createScheduler()
	})
	return schedulerInstance
}

func createScheduler() scheduler.Scheduler {
	panic(wire.Build(
		scheduler.New,
	))
}
