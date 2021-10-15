package config

import "os"

type DBConfig struct {
	Host     string
	Port     string
	DBName   string
	User     string
	Password string
	SSLMode  string
	Schema   string
}

var databaseConfig DBConfig

func init() {
	databaseConfig.Host = getEnvDefault("POSTGRES_HOST", "localhost")
	databaseConfig.Port = getEnvDefault("POSTGRES_PORT", "5432")
	databaseConfig.User = getEnvDefault("POSTGRES_USER", "matcha")
	databaseConfig.Password = getEnvDefault("POSTGRES_PASSWORD", "")
	databaseConfig.DBName = getEnvDefault("POSTGRES_DB", "matcha_db")
	databaseConfig.SSLMode = getEnvDefault("POSTGRES_SSL", "")
	databaseConfig.Schema = getEnvDefault("POSTGRES_SCHEMA", databaseConfig.User)
}

// Get database configurations from enviroment variables
func GetDBConfig() DBConfig {
	return databaseConfig
}

// Get database source name string, config is taken with `GetDBConfig()`
func GetDSN() string {
	dsnvals := map[string]string{
		"host":     databaseConfig.Host,
		"port":     databaseConfig.Port,
		"dbname":   databaseConfig.DBName,
		"user":     databaseConfig.User,
		"password": databaseConfig.Password,
	}

	if len(databaseConfig.SSLMode) != 0 {
		dsnvals["sslmode"] = databaseConfig.SSLMode
	}

	var sep, dsnstr string
	for k, v := range dsnvals {
		dsnstr += sep + k + "=" + v
		sep = " "
	}

	return dsnstr
}

// Get database driver name string
func GetDBDriverName() string {
	return "postgres"
}

// Returns a value of enviroment variable with the name `varname`.
// If variable isn't set, returns `deflt_value`
func getEnvDefault(varname string, deflt_value string) string {
	varvalue := os.Getenv(varname)
	if len(varvalue) == 0 {
		return deflt_value
	}
	return varvalue
}
