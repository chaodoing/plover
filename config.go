package plover

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
}

func NewConfig() *Config {
	config := &Config{viper.New()}
	err := config.InitConfig()
	if err != nil {
		log.Printf("[Error] 错误: %s", err.Error())
	}
	return config
}
func (c *Config) InitConfig() error {
	c.SetConfigName("app")
	c.AddConfigPath("conf")
	c.SetConfigType("yaml")
	return c.ReadInConfig()
}

func (c *Config) GetDataDefault(data interface{}, defaults interface{}) interface{} {
	if data != nil {
		return data
	}
	return defaults
}
