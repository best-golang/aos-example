package env

import (
	"os"

	"github.com/spf13/viper"

	"github.com/aos-stack/env"
)

type EnvExampleCommand struct{}

func (c EnvExampleCommand) Execute() {
	nowEnv := os.Getenv("SYSTEM_ENV")
	if nowEnv != "" {
		viper.Set("application.enviroment", nowEnv)
	}
	env.SetProvider(EnvExampleCommand{})
}

func (e EnvExampleCommand) Get() int {
	getEnvVal := env.Local
	switch viper.GetString("application.enviroment") {
	case "test":
		getEnvVal = env.Testing
	case "prepare":
		getEnvVal = env.Staging
	case "production":
		getEnvVal = env.Production
	}
	return getEnvVal
}

func (e EnvExampleCommand) GetLabel() string {
	return viper.GetString("application.enviroment")
}
