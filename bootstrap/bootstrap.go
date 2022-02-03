package bootstrap

import (
	"github.com/vsyakunin/sg-task/domain/models/config"

	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

func InitConfig(path string) *config.Config {
	var cfg config.Config
	_, err := toml.DecodeFile(path, &cfg)
	if err != nil {
		log.Fatalln(err)
	}

	return &cfg
}
