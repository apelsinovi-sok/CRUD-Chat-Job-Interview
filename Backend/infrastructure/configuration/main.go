package configuration

import "github.com/spf13/viper"

type Config struct {
	DatabaseUrl string `mapstructure:"DATABASE_URL"`
	Port        string `mapstructure:"PORT"`
}

func New(path string) (config Config) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err.Error())
	}
	//configuration.ParseDbUrl()
	viper.WatchConfig()
	return
}
