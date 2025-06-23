package configs

import "github.com/spf13/viper"

type Config struct {
	DBDriver          string `mapstructure:"db_driver"`
	DBURL             string `mapstructure:"db_url"`
	RabbitMQURL       string `mapstructure:"rabbitmq_url"`
	WebServerPort     string `mapstructure:"web_server_port"`
	OrderCreatedEvent string `mapstructure:"order_created_event"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg *Config
	viper.SetConfigName("configs")
	viper.SetConfigType("yml")
	viper.AddConfigPath(path)
	viper.SetEnvPrefix("viper")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
