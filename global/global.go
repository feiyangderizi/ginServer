package global

import (
	"github.com/sadlil/gologger"

	"github.com/feiyangderizi/ginServer/initialize/config"
)

var (
	Config *config.ServerConfig
	Logger gologger.GoLogger
)
