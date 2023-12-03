package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	// debug, info, warn, error,
	LogLevel         string
	HttpPort         string
	GrpcPort         string
	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
	SmsProvideApiKey string
	JWTSecret        string
	JWTExpiresInSec  int
}

func Load() (Config, error) {
	v := viper.New()
	v.AutomaticEnv()

	var config Config
	v.SetDefault("LOG_LEVEL", "debug")
	v.SetDefault("HTTP_PORT", ":3030")
	v.SetDefault("GRPC_PORT", ":5000")
	v.SetDefault("POSTGRES_HOST", "localhost")
	v.SetDefault("POSTGRES_PORT", 5432)
	v.SetDefault("POSTGRES_USER", "postgres")
	v.SetDefault("POSTGRES_PASSWORD", "postgres")
	v.SetDefault("POSTGRES_DATABASE", "postgres")
	v.SetDefault("SMS_PROVIDER_API_KEY", "")
	v.SetDefault("JWT_SECRET", "")
	v.SetDefault("JWT_EXPIRES_IN_SEC", 2630000) // 1 month

	config.LogLevel = v.GetString("LOG_LEVEL")
	config.HttpPort = v.GetString("HTTP_PORT")
	config.GrpcPort = v.GetString("GRPC_PORT")
	config.PostgresHost = v.GetString("POSTGRES_HOST")
	config.PostgresPort = v.GetInt("POSTGRES_PORT")
	config.PostgresUser = v.GetString("POSTGRES_USER")
	config.PostgresPassword = v.GetString("POSTGRES_PASSWORD")
	config.PostgresDatabase = v.GetString("POSTGRES_DATABASE")
	config.SmsProvideApiKey = v.GetString("SMS_PROVIDER_API_KEY")
	config.JWTSecret = v.GetString("JWT_SECRET")
	config.JWTExpiresInSec = v.GetInt("JWT_EXPIRES_IN_SEC")

	return config, nil
}

func (c Config) NewLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	level, err := zapcore.ParseLevel(c.LogLevel)
	if err != nil {
		return nil, err
	}

	config.Level.SetLevel(level)
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
