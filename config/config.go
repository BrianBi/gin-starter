package config

import (
	"fmt"
	"github.com/brianbi/gin-starter/pkg/helpers"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cast"
	"github.com/spf13/viper"

	"github.com/brianbi/gin-starter/config/configs"
	"github.com/brianbi/gin-starter/pkg/file"
)

// Configuration configuration
type Configuration struct {
	App 			configs.App 			`mapstructure:"app"`
	Database 		configs.Database 		`mapstructure:"database"`
	FileSystems 	configs.FileSystems 	`mapstructure:"filesystems"`
	Log 			configs.Log 			`mapstructure:"log"`
	Mail 			configs.Mail 			`mapstructure:"mail"`
}

var (
	v 			= viper.New()
	config 		= new(Configuration)
	configPath 	= "./config/"
	configType 	= "yaml"
	configName  = "config"
)

func Initialize(filePath string) Configuration {
	if len(filePath) == 0 {
		v.SetConfigName(configName)
		v.SetConfigType(configType)
		v.AddConfigPath(configPath)
		filePath = configPath + configName + "." + configType
	} else {
		v.SetConfigFile(filePath)
	}

	if _, ok := file.IsExists(filePath); !ok {
		panic("config file: " + filePath + " not found")
	}

	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		if err := v.Unmarshal(&config); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&config); err != nil {
		panic(err)
	}

	return *config
}

// GetConfig struct
func GetConfig() Configuration {
	return *config
}

// get current package value by path
func getValue(path string, value ...interface{}) interface{} {
	if !v.IsSet(path) || helpers.Empty(v.Get(path)) {
		if len(value) > 0 {
			return value[0]
		}

		return nil
	}

	return v.Get(path)
}

// Get get config value to string
func Get(path string, value ...interface{}) string {
	return GetString(path, value...)
}

// GetString Obtain the configuration information of type string
func GetString(path string, value ...interface{}) string {
	return cast.ToString(getValue(path, value...))
}

// GetInt Obtain the configuration information of type int
func GetInt(path string, value ...interface{}) int {
	return cast.ToInt(getValue(path, value...))
}

// GetFloat64 Obtain the configuration information of type Float64
func GetFloat64(path string, value ...interface{}) float64 {
	return cast.ToFloat64(getValue(path, value...))
}

// GetInt64 Obtain the configuration information of type Int64
func GetInt64(path string, value ...interface{}) int64 {
	return cast.ToInt64(getValue(path, value...))
}

// GetUint Obtain the configuration information of type Uint
func GetUint(path string, value ...interface{}) uint {
	return cast.ToUint(getValue(path, value...))
}

// GetBool Obtain the configuration information of type Bool
func GetBool(path string, value ...interface{}) bool {
	return cast.ToBool(getValue(path, value...))
}

// GetStringMapString get structural data
func GetStringMapString(path string) map[string]string {
	return v.GetStringMapString(path)
}
