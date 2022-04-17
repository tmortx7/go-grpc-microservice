package config

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type config struct {
	Database DatabaseConfig
	Server   ServerConfig
}
type DatabaseConfig struct {
	User                 string
	Password             string
	Net                  string
	Addr                 string
	DBName               string
	AllowNativePasswords bool
	Params               struct {
		ParseTime string
		Charset   string
		Loc       string
	}
}
type ServerConfig struct {
	Address string
}

// C is config variable
var C config

// ReadConfig configures config file
func ReadConfig() {
	Config := &C
	viper.AddConfigPath(filepath.Join(rootDir(), "config"))
	viper.SetConfigName("config")

	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
