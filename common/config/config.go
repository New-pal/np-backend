package config

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const envPrefix = "newpal"

func getFileName() string {
	fileName := "config"
	env := os.Getenv(strings.ToUpper(envPrefix + "_env"))
	if env == "" {
		env = "dev"
	}
	return fileName + "." + env
}

func getBaseDir() string {
	_, b, _, ok := runtime.Caller(0)
	if !ok {
		panic("config: unable to determine the caller.")
	}
	return filepath.Dir(b)
}

func setPaths(baseDir string) {
	viper.AddConfigPath(".")
	viper.AddConfigPath(baseDir + "/yml")
	viper.AddConfigPath("/etc/" + envPrefix)
	viper.AddConfigPath("$HOME/." + envPrefix)
}

func Setup() {
	baseDir := getBaseDir()
	viper.SetConfigName(getFileName())
	setPaths(baseDir)
	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic("viper err " + err.Error())
	}
}
