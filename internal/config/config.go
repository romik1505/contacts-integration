package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

type ConfigFile struct {
	DBConfig         DBConfig
	APISecretKey     string
	HostUrl          string
	BeanstalkdConfig BeanstalkdConfig
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

type BeanstalkdConfig struct {
	Host string
	Port string
}

func (b BeanstalkdConfig) ConnectionString() string {
	return fmt.Sprintf("%s:%s", b.Host, b.Port)
}

var (
	Config   ConfigFile
	envLevel string
)

var once sync.Once

func init() {
	once.Do(func() {
		envLevel = os.Getenv("ENVIRONMENT")
		if envLevel == "local" {
			log.Printf("CONFIG env level = %s\n", envLevel)
		}

		conf, err := NewConfig()
		if err != nil {
			log.Fatalf("config initiation error: %v", err)
		}
		Config = conf
	})
}

func NewDBConfig() (DBConfig, error) {
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

func NewBeanstalkdConfig() (BeanstalkdConfig, error) {
	host, err := GetEnv("BEANSTALKD_HOST")
	if err != nil {
		return BeanstalkdConfig{}, err
	}
	port, err := GetEnv("BEANSTALKD_PORT")
	if err != nil {
		return BeanstalkdConfig{}, err
	}
	return BeanstalkdConfig{
		Host: host,
		Port: port,
	}, nil
}

func NewConfig() (ConfigFile, error) {
	var envFile string
	envFile, ok := os.LookupEnv("ENV_FILE")
	if !ok {
		envFile = ".env"
	}
	log.Printf("env file: %s", envFile)

	if err := godotenv.Load(envFile); err != nil {
		log.Printf("%s not found: %v\n", envFile, err)
	}

	dbConfig, err := NewDBConfig()
	if err != nil {
		return ConfigFile{}, err
	}

	bsConfig, err := NewBeanstalkdConfig()
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
		DBConfig:         dbConfig,
		APISecretKey:     apiSecret,
		HostUrl:          hostUrl,
		BeanstalkdConfig: bsConfig,
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
