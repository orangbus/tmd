package bootstrap

import (
	"github.com/orangbus/cmd/pkg/config"
	"github.com/orangbus/cmd/pkg/search"
)

func SetUp() {
	config.LoadConfig()
	SetupDatabase()
	search.NewSearch()
}
