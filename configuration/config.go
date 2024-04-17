package configuration

import (
	"path"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

const (
	Prod string = "prod"
	Dev  string = "dev"
	Test string = "test"
)

const (
	Debug string = "debug"
	Info  string = "info"
	Warn  string = "warn"
	Error string = "error"
)

type Config struct {
	App struct {
		Env        string
		Debug      bool
		Pagination struct {
			MaxSize int `mapstructure:"max_size"`
		}
	}
	Server struct {
		Port string
	}
	Database struct {
		Name       string
		Host       string
		Port       int
		SSL        string
		User       string
		Password   string
		Timezone   string
		Connection struct {
			Idle int
			Open int
		}
	}
	Messaging struct {
		Protocol string
		Host     string
		Port     int
		User     string
		Password string
		Queue    struct {
			Name string
		}
	}
	Logging struct {
		Level string
	}
}

var config *Config

func setup() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigFile(dir + "/config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
}

func NewConfig() Config {
	if config == nil {
		setup()
	}
	return *config
}
