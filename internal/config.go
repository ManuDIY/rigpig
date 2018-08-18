package internal

import "github.com/spf13/viper"

var Logfile = "rigpig.log"

var NotDone = true

func LoadConfig() (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigFile("rigpig")
	v.SetConfigType("YAML")
	v.AddConfigPath(".")
	v.AddConfigPath("/etc/rigpig")
	v.AddConfigPath("$HOME/.rigpig")
	v.AutomaticEnv()
	err := v.ReadInConfig()
	return v, err
}
