package config

import (
	"fmt"
	"os"
)

// Returns a value of enviroment variable with the name `varname`.
// If variable isn't set, returns `deflt_value`
func getEnvDefault(varname string, deflt_value string) string {
	varvalue := os.Getenv(varname)
	if len(varvalue) == 0 {
		return deflt_value
	}
	return varvalue
}

type DBConfig struct {
	Host     string
	Port     string
	DBName   string
	User     string
	Password string
	SSLMode  string
}

// Get database configurations from enviroment variables
func GetDBConfig() DBConfig {
	return DBConfig{
		Host:     getEnvDefault("POSTGRES_HOST", "localhost"),
		Port:     getEnvDefault("POSTGRES_PORT", "5432"),
		User:     getEnvDefault("POSTGRES_USER", ""),
		Password: getEnvDefault("POSTGRES_PASSWORD", ""),
		DBName:   getEnvDefault("POSTGRES_DB", ""),
		SSLMode:  getEnvDefault("POSTGRES_SSL", "disable"),
	}
}

// Get database source name string, config is taken with `GetDBConfig()`
func GetDSN() string {
	dbconfig := GetDBConfig()
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		dbconfig.Host,
		dbconfig.Port,
		dbconfig.DBName,
		dbconfig.User,
		dbconfig.Password,
		dbconfig.SSLMode)
}
