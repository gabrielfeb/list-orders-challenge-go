package configs

import (
	"[github.com/spf13/viper](https://github.com/spf13/viper)"
)

// Config define todas as configurações necessárias para a aplicação.
// As tags `mapstructure` são usadas pelo Viper para decodificar os valores
// do arquivo de configuração (configs.yml) para esta struct.
type Config struct {
	DBDriver          string `mapstructure:"db_driver"`
	DBURL             string `mapstructure:"db_url"`
	RabbitMQURL       string `mapstructure:"rabbitmq_url"`
	WebServerPort     string `mapstructure:"web_server_port"`
	GRPCServerPort    string `mapstructure:"grpc_server_port"`
	OrderCreatedEvent string `mapstructure:"order_created_event"`
}

// LoadConfig carrega as configurações a partir de um arquivo .yml no caminho especificado.
// Ele também utiliza o Viper para permitir que variáveis de ambiente sobrescrevam
// os valores do arquivo, seguindo as melhores práticas de 12-Factor App.
func LoadConfig(path string) (*Config, error) {
	var cfg *Config
	viper.SetConfigName("configs")
	viper.SetConfigType("yml")
	viper.AddConfigPath(path)
	viper.SetEnvPrefix("viper")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
