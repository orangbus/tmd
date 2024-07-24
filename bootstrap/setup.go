package bootstrap

import (
	"github.com/orangbus/cmd/pkg/config"
)

func SetUp() {
	config.LoadConfig()
	//SetupDatabase()
	//search.NewSearch()
}
