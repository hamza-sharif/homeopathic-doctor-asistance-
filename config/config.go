// Package config provides configuration settings for the application.
package config

import (
	"github.com/spf13/viper"
)

// keys for database configuration.
const (
	APIKey     = "api.key"
	DBName     = "auth.db.name"
	DBHost     = "auth.db.ip"
	DBPort     = "auth.db.port"
	DBUser     = "auth.db.user"
	DBPassword = "auth.db.password"

	ServerHost     = "server.host"
	ServerPort     = "server.port"
	ServerEmail    = "server.email"
	ServerPassword = "server.password"

	ErrorMessageToken401 = "error.token.401"
	ErrorMessageJSON422  = "error.token.422"
	ErrorMessage500      = "error.500"
	ErrorMessage400      = "error.400"
	ErrorMessage404      = "error.404"
	ErrorMessage409      = "error.409"
	ErrorEmailNotFound   = "email.notfound"

	LogLevel = "log.level"
)

func init() {
	// env var for db.
	_ = viper.BindEnv(APIKey, "API_KEY")
	_ = viper.BindEnv(DBName, "AUTH_DB_NAME")
	_ = viper.BindEnv(DBHost, "AUTH_DB_HOST")
	_ = viper.BindEnv(DBUser, "AUTH_DB_USER")
	_ = viper.BindEnv(DBPassword, "AUTH_DB_PASSWORD")
	_ = viper.BindEnv(DBPort, "AUTH_DB_PORT")

	_ = viper.BindEnv(ServerHost, "SERVER_HOST")
	_ = viper.BindEnv(ServerPort, "SERVER_PORT")

	_ = viper.BindEnv(LogLevel, "LOG_LEVEL")

	// defaults.
	defaultVariableValues()

	// error message.
	defaultErrorMessages()
}

func defaultVariableValues() {
	viper.SetDefault(DBName, "hospital-datastore")
	viper.SetDefault(DBHost, "0.0.0.0")
	viper.SetDefault(DBPort, "5432")
	viper.SetDefault(DBUser, "test")
	viper.SetDefault(DBPassword, "myproject123")

	viper.SetDefault(ServerHost, "0.0.0.0")
	viper.SetDefault(ServerPort, "8080")
	viper.SetDefault(ServerEmail, "noreply@wanclouds.net")
	viper.SetDefault(ServerPassword, "6_CfHuNwVIjZYClUYCDxgQ")
	viper.SetDefault(LogLevel, "debug")

}

func defaultErrorMessages() {
	viper.SetDefault(ErrorMessageToken401, "failed to verify token")
	viper.SetDefault(ErrorMessageJSON422, "Password  not match")
	viper.SetDefault(ErrorMessage500, "Sorry! Something went wrong on making request")
	viper.SetDefault(ErrorMessage404, "Unable to find the resource")
	viper.SetDefault(ErrorMessage409, "requested resource is already Modify")
	viper.SetDefault(ErrorEmailNotFound, "provided email does not exist")
}
