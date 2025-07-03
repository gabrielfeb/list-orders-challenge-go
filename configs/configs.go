package configs

import (
	"os"

	"github.com/spf13/viper"
)

type Conf struct{ DBURL, GraphQLPort, GRPCServerPort string }

func LoadConfig(path string) (*Conf, error) {
	viper.SetConfigName("configs")
	viper.AddConfigPath(path)
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	return &Conf{
		DBURL:          os.Getenv("DB_URL"),
		GraphQLPort:    os.Getenv("GRAPHQL_PORT"),
		GRPCServerPort: os.Getenv("GRPC_PORT"),
	}, nil
}
