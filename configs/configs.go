package configs

import (
	"github.com/spf13/viper"
)

/*
Constants is a struct which store the const within the projsct
*/
type Constants struct {
	Host          string
	Port          string
	Portal        Portal
	CoureseSystem CoureseSystem
}

/*
Portal contains all the urls of application
*/
type Portal struct {
	Base                 string
	IndexPage            string
	Login                string
	MainPage             string
	AptreeListPage       string
	AptreeAAListPage     string
	SsoLoginCourseSystem string
}

/*
CoureseSystem contains all of the elements from course system
*/
type CoureseSystem struct {
	Base     string
	MainPage string
	Select   string
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
