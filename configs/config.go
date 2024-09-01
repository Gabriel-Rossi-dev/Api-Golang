package configs

import "github.com/spf13/viper"

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	DataBase string
}

func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432") //porta padrao PostGres
}

func load() error {
	// Nome do arquivo de configuração
	viper.SetConfigName("config")
	//tipo do arquivo
	viper.SetConfigType("toml")
	// o '.' representa que o arquivo será buscado na mesma pasta que contem o arquivo binário
	viper.AddConfigPath(".")
	//variável recebe o que o pacote viper leu do arquivo
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	//cria um ponteiro da config
	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		DataBase: viper.GetString("database.name"),
	}

	return nil

}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}
