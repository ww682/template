package config

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GlobalMysqlConn *gorm.DB
var GlobalRedisConn *redis.Client

type Config struct {
	Mysql MysqlConfig `yaml:"mysql"`
	Redis RedisConfig `yaml:"redis"`
}

type MysqlConfig struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Database     string `yaml:"database"`
	Charset      string `yaml:"charset"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	MaxOpenConns int    `yaml:"max_open_conns"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
}

func LoadConfig(path string) (*Config, error) {
	cfg := &Config{}
	config, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal([]byte(config), cfg)
	return cfg, err
}

func InitMysql(cfg *MysqlConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.Charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	GlobalMysqlConn = db
	return db, nil
}
func InitRedis(cfg *RedisConfig) (*redis.Client, error) {

	GlobalRedisConn = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.Database,
	})
	return GlobalRedisConn, nil
}
