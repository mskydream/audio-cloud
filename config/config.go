package config

import "github.com/spf13/viper"

type Config struct {
	Database DB     `mapstructure:"db"`
	SaveDir  string `mapstructure:"savedir"`
	Auth     Auth   `mapstructure:"auth"`
}

type DB struct {
	Address  string `mapstructure:"address"`
	User     string `mapstructure:"user"`
	Password string `mapstucture:"password"`
	Name     string `mapstructure:"name"`
}
type Auth struct {
	AccessTokenTTL  string `mapstructure:"accessTokenTTL"`
	RefreshTokenTTL string `mapstructure:"refreshTokenTTL"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
