package config

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

const (
	defaultHttpPort               = "8020"
	defaultHttpRWTimeout          = 10 * time.Second
	defaultHttpMaxHeaderMegabytes = 1

	EnvLocal = "local"
	Prod     = "prod"
)

type (
	Config struct {
		Environment string
		HTTP        HTTPConfig
		DB          DBConfig
		Kafka       KafkaConfig
		ApiToken    string
	}

	HTTPConfig struct {
		Host               string        `mapstructure:"host"`
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
	}
	DBConfig struct {
		Host     string `mapstructure:"DB_HOSTNAME"`
		Port     string `mapstructure:"DB_PORT"`
		Username string `mapstructure:"DB_USERNAME"`
		Password string `mapstructure:"DB_PASSWORD"`
		DBName   string `mapstructure:"DB_NAME"`
	}
	KafkaConfig struct {
		Host string `mapstructure:"kafka_host"`
		Port string `mapstructure:"kafka_port"`
	}
)

func Init(configDir string) (*Config, error) {
	populateDefaults()

	if err := parseEnv(); err != nil {
		return nil, err
	}
	if err := parseConfigFile(configDir, viper.GetString("env")); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}
	/* to-do set from .env */
	setFromEnv(&cfg)
	return &cfg, nil

}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("db", &cfg.DB); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("kafka", &cfg.Kafka); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("api_token", &cfg.ApiToken); err != nil {
		return err
	}

	return nil
}

func parseConfigFile(folder, env string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if env == EnvLocal {
		return nil
	}

	viper.SetConfigName(env)

	return viper.MergeInConfig()
}

func setFromEnv(cfg *Config) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	cfg.DB.Host = viper.GetString("DB_HOSTNAME")
	cfg.DB.Port = viper.GetString("DB_PORT")
	cfg.DB.Username = viper.GetString("DB_USERNAME")
	cfg.DB.Password = viper.GetString("DB_PASSWORD")
	cfg.DB.DBName = viper.GetString("DB_NAME")
}

func populateDefaults() {
	viper.SetDefault("http.port", defaultHttpPort)
	viper.SetDefault("http.max_header_megabytes", defaultHttpMaxHeaderMegabytes)
	viper.SetDefault("http.timeouts.read", defaultHttpRWTimeout)
	viper.SetDefault("http.timeouts.write", defaultHttpRWTimeout)
}

func parseEnv() error {
	if err := parseHostFromEnv(); err != nil {
		return err
	}

	if err := parsingDBFromEnv(); err != nil {
		return err
	}

	return parsePasswordFromEnv()
}

func parseHostFromEnv() error {
	viper.SetEnvPrefix("http")
	return viper.BindEnv("hostname")
}

func parsePasswordFromEnv() error {
	viper.SetEnvPrefix("password")

	return viper.BindEnv("salt")
}

func parsingDBFromEnv() error {
	viper.SetEnvPrefix("db")

	if err := viper.BindEnv("hostname"); err != nil {
		return err
	}
	if err := viper.BindEnv("port"); err != nil {
		return err
	}
	if err := viper.BindEnv("username"); err != nil {
		return err
	}
	if err := viper.BindEnv("name"); err != nil {
		return err
	}

	return viper.BindEnv("password")
}
