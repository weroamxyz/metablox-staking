package settings

import (
	"runtime"

	"github.com/fsnotify/fsnotify"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Init() (err error) {
	_, fileName, _, _ := runtime.Caller(0)
	filePath := fileName[:len(fileName)-20]
	viper.SetConfigFile(filePath + "/config.yaml")
	err = viper.ReadInConfig()

	profile := viper.GetString("profile")

	if profile != "" {
		viper.SetConfigFile(filePath + "/config-" + profile + ".yaml")
		viper.MergeInConfig()
	}

	if err != nil {
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		logger.Info("config file has been changed")
	})
	return nil
}
