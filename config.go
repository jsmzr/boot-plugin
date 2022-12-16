package plugin

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const configPrefix string = "boot.config"

var defaultConfig map[string]interface{} = map[string]interface{}{"file": "application.yaml", "path": "."}

func initEnv() {
	// env start with BOOT_
	viper.SetEnvPrefix("BOOT")
	viper.AutomaticEnv()

}

func initDefault() {
	// default config file, if you don't set by envs and flags, use this
	for key := range defaultConfig {
		viper.SetDefault(configPrefix+"."+key, defaultConfig[key])
	}
}

func initLocalConfig() error {
	// if file contains suffix, then split
	file := viper.GetString(configPrefix + ".file")
	viper.SetConfigFile(file)
	viper.AddConfigPath(viper.GetString(configPrefix + ".path"))
	if err := viper.ReadInConfig(); err != nil {
		if e, ok := err.(*os.PathError); ok {
			log(fmt.Sprintf("read config failed, %v", e))
		} else {
			return err
		}
	}
	return nil
}

func initConfig() error {
	initEnv()
	initDefault()
	return initLocalConfig()
}
