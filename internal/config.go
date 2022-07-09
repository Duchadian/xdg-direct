package config

import "github.com/spf13/viper"

type Mapping struct {
	Url     string   `mapstructure:"url"`
	Command string   `mapstructure:"cmd"`
	Args    []string `mapstructure:"args"`
}

func ReadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/xdg-direct")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func UrlMappings() []Mapping {
	var mappings []Mapping
	if err := viper.UnmarshalKey("mappings", &mappings); err != nil {
		panic(err)
	}

	return mappings
}

func DefaultCommand() string {
	return viper.GetString("default_command")
}

func DebugMode() bool {
	return viper.GetBool("debug_mode")
}
