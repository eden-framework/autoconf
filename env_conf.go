package autoconf

import (
	"github.com/profzone/envconfig"
	"github.com/sirupsen/logrus"
)

func FromEnv(confPrefix string, conf []interface{}) {
	// atempt to load local.yml
	_ = envconfig.LoadDefaultFromYml("./build/configs/local.yml")

	for _, c := range conf {
		err := envconfig.Process(confPrefix, c)
		if err != nil {
			logrus.Panic(err)
		}
		envconfig.Usage(confPrefix, c, true)
	}
}

func GetDefaultFromEnv(confPrefix string, conf []interface{}) []envconfig.EnvVar {
	var defaultEnvVars = make([]envconfig.EnvVar, 0)
	for _, c := range conf {
		envs, err := envconfig.GatherInfo(confPrefix, c)
		if err != nil {
			logrus.Panic(err)
		}
		defaultEnvVars = append(defaultEnvVars, envs...)
	}
	return defaultEnvVars
}
