package config

import "fmt"

const (
	DBUser     = "postgres"
	DBPassword = "postgres00"
	DBName     = "escola_biblica"
	DBHost     = "localhost"
	DBPort     = "5435"
	DBType     = "postgres"
)

// GetDBType get type for database connection
func GetDBType() string {
	return DBType
}

// GetPostgresConnectionString get config connections database
func GetPostgresConnectionString() string {
	dataBase := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=America/Fortaleza",
		DBHost,
		DBPort,
		DBUser,
		DBName,
		DBPassword)

	return dataBase
}
