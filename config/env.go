package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file")
		return
	}
	log.Println(".env file loaded")
}

// Функция получения значения по ключу из переменных окружения
func getString(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// Функция получения целого значения по ключу из переменых окружения
func getInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	i, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return i
}

// Функция получения булевой переменной из переменных окружения
func getBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	b, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return b
}

// Структура для конфигурации доступа к базе данных
type DatabaseConfig struct {
	url string
}

// Конструктор конфигурации базы данных
func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		url: getString("DATABASE_URL", ""),
	}
}

// Структура для конфигурации логирования приложения
type LogConfig struct {
	Level  int
	Format string
}

// Конструктор конфигурации логирования
func NewLogCofig() *LogConfig {
	return &LogConfig{
		Level:  getInt("LOG_LEVEL", 0),
		Format: getString("LOG_FORMAT", "json"),
	}
}
