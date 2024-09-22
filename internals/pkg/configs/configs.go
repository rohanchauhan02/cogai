package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type ImmutableConfig interface {
	GetOpenAIKey() string
}

type config struct {
	openAIKey string `mapstructure:"OPENAI_API_KEY"`
}

func (c *config) GetOpenAIKey() string {
	return c.openAIKey
}

var (
	im *config
	// imOnce *sync.Once
)

func NewImmutableConfig() ImmutableConfig {
	// imOnce.Do(func() {
	v := viper.New()
	v.SetConfigName("app.config.yml")
	v.AddConfigPath(".")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		fmt.Println(500, err, "[CONFIG][missing] Failed to read app.config.yml file", "failed")
		panic(err)
	}
	v.Unmarshal(im)
	// })
	return im
}
