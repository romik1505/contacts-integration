package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type ConfigFile struct {
	DBConfig     DBConfig
	APISecretKey string
	HostUrl      string
}

type DBConfig struct {
	Host     string
	Driver   string
	User     string
	Password string
	Name     string
	Port     string
}

func (d DBConfig) ConnectionString() string {
	return fmt.Sprintf(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.Name,
	))
}

var (
	Config   ConfigFile
	envLevel string
)

func init() {
	envLevel = os.Getenv("ENVIRONMENT")
	if envLevel == "local" {
		log.Printf("CONFIG env level = %s\n", envLevel)
	}

	conf, err := NewConfig()
	if err != nil {
		log.Fatalf("config initiation error: %v", err)
	}
	Config = conf
}

func NewDBConfig() (DBConfig, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf(".env not found: %v\n", err)
	}

	DBHost, err := GetEnv("DB_HOST")
	if err != nil {
		return DBConfig{}, err
	}

	DBDriver, err := GetEnv("DB_DRIVER")
	if err != nil {
		return DBConfig{}, err
	}

	DBUser, err := GetEnv("DB_USER")
	if err != nil {
		return DBConfig{}, err
	}

	DBPassword, err := GetEnv("DB_PASSWORD")
	if err != nil {
		return DBConfig{}, err
	}

	DBName, err := GetEnv("DB_NAME")
	if err != nil {
		return DBConfig{}, err
	}

	DBPort, err := GetEnv("DB_PORT")
	if err != nil {
		return DBConfig{}, err
	}

	return DBConfig{
		Host:     DBHost,
		Driver:   DBDriver,
		User:     DBUser,
		Password: DBPassword,
		Name:     DBName,
		Port:     DBPort,
	}, nil
}

func NewConfig() (ConfigFile, error) {
	dbConfig, err := NewDBConfig()
	if err != nil {
		return ConfigFile{}, err
	}

	apiSecret, err := GetEnv("API_SECRET")
	if err != nil {
		return ConfigFile{}, err
	}

	hostUrl, err := GetEnv("HOST_URL")
	if err != nil {
		return ConfigFile{}, err
	}

	return ConfigFile{
		DBConfig:     dbConfig,
		APISecretKey: apiSecret,
		HostUrl:      hostUrl,
	}, nil
}

func GetEnv(key string) (string, error) {
	if envLevel == "local" {
		key = "TEST_" + key
	}
	value, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("%s not set", key)
	}

	log.Printf("CONFIG: %s=%s", key, value)
	return value, nil
}
