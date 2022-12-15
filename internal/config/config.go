package config

import (
	"github.com/spf13/viper"
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
	}

	HTTPConfig struct {
		Host               string        `mapstructure:"host"`
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
	}
	DBConfig struct {
		Host     string `mapstructure:"hostname"`
		Port     string `mapstructure:"port"`
		Username string `mapstructure:"userName"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"name"`
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
	//setFromEnv(&cfg)
	return &cfg, nil

}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("db", &cfg.DB); err != nil {
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
	cfg.HTTP.Host = viper.GetString("host")
	cfg.DB.Host = viper.GetString("hostname")
	cfg.DB.Port = viper.GetString("port")
	cfg.DB.Username = viper.GetString("username")
	cfg.DB.Password = viper.GetString("password")
	cfg.DB.DBName = viper.GetString("name")
	cfg.Environment = viper.GetString("env")
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
