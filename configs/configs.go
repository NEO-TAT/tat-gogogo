package configs

import (
	"github.com/spf13/viper"
)

/*
Constants is a struct which store the const within the projsct
*/
type Constants struct {
	HOST   string
	PORT   string
	PORTAL Portal
}

/*
Portal contains all the urls of application
*/
type Portal struct {
	BASE                 string
	IndexPage            string
	Login                string
	MainPage             string
	AptreeListPage       string
	AptreeAAListPage     string
	SsoLoginCourseSystem string
}

/*
Config is a struct which store project settings
*/
type Config struct {
	Constants
}

/*
New is used to create the new configuration of project
*/
func New() (*Config, error) {
	config := Config{}
	constants, err := initViper()
	config.Constants = constants

	return &config, err
}

func initViper() (Constants, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")

	err := viper.ReadInConfig()
	if err != nil {
		return Constants{}, err
	}

	var constants Constants
	err = viper.Unmarshal(&constants)

	return constants, err
}
