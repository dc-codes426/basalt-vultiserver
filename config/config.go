package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port           int64  `mapstructure:"port" json:"port,omitempty"`
		Host           string `mapstructure:"host" json:"host,omitempty"`
		VaultsFilePath string `mapstructure:"vaults_file_path" json:"vaults_file_path,omitempty"`
	} `mapstructure:"server" json:"server"`

	Redis struct {
		Host     string `mapstructure:"host" json:"host,omitempty"`
		Port     string `mapstructure:"port" json:"port,omitempty"`
		User     string `mapstructure:"user" json:"user,omitempty"`
		Password string `mapstructure:"password" json:"password,omitempty"`
		DB       int    `mapstructure:"db" json:"db,omitempty"`
	} `mapstructure:"redis" json:"redis,omitempty"`

	Relay struct {
		Server string `mapstructure:"server" json:"server"`
	} `mapstructure:"relay" json:"relay,omitempty"`

	BlockStorage struct {
		Host      string `mapstructure:"host" json:"host"`
		Region    string `mapstructure:"region" json:"region"`
		AccessKey string `mapstructure:"access_key" json:"access_key"`
		SecretKey string `mapstructure:"secret" json:"secret"`
		Bucket    string `mapstructure:"bucket" json:"bucket"`
	} `mapstructure:"block_storage" json:"block_storage"`
}

func GetConfigure() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.host", "localhost")
	viper.SetDefault("server.vaults_file_path", "vaults")
	viper.SetDefault("redis.host", "localhost")
	viper.SetDefault("redis.port", "6379")
	viper.SetDefault("redis.user", "")
	viper.SetDefault("redis.password", "")
	viper.SetDefault("redis.db", 0)
	viper.SetDefault("relay.server", "https://api.vultisig.com/router")

	// Bind env vars so Unmarshal picks them up without a config file.
	_ = viper.BindEnv("server.port", "SERVER_PORT")
	_ = viper.BindEnv("server.host", "SERVER_HOST")
	_ = viper.BindEnv("server.vaults_file_path", "SERVER_VAULTS_PATH")
	_ = viper.BindEnv("redis.host", "REDIS_HOST")
	_ = viper.BindEnv("redis.port", "REDIS_PORT")
	_ = viper.BindEnv("redis.user", "REDIS_USER")
	_ = viper.BindEnv("redis.password", "REDIS_PASSWORD")
	_ = viper.BindEnv("redis.db", "REDIS_DB")
	_ = viper.BindEnv("relay.server", "RELAY_SERVER")
	_ = viper.BindEnv("block_storage.host", "BLOCK_STORAGE_HOST")
	_ = viper.BindEnv("block_storage.region", "BLOCK_STORAGE_REGION")
	_ = viper.BindEnv("block_storage.access_key", "BLOCK_STORAGE_ACCESS_KEY")
	_ = viper.BindEnv("block_storage.secret", "BLOCK_STORAGE_SECRET")
	_ = viper.BindEnv("block_storage.bucket", "BLOCK_STORAGE_BUCKET")

	// Config file is optional — env vars alone are sufficient.
	_ = viper.ReadInConfig()

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
