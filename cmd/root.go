package cmd

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
	"github.com/wiliehidayat87/rmqp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "A generator for Cobra based Applications",
		Long:  `Cobra is a CLI library for Go that empowers applications.`,
	}
)

var (
	APP_URL     string = getEnv("APP_URL")
	APP_PORT    string = getEnv("APP_PORT")
	APP_TZ      string = getEnv("APP_TZ")
	API_VERSION string = getEnv("API_VERSION")
	URI_PGSQL   string = getEnv("URI_PGSQL")
	URI_REDIS   string = getEnv("URI_REDIS")
	RMQ_HOST    string = getEnv("RMQ_HOST")
	RMQ_USER    string = getEnv("RMQ_USER")
	RMQ_PASS    string = getEnv("RMQ_PASS")
	RMQ_PORT    string = getEnv("RMQ_PORT")
	RMQ_VHOST   string = getEnv("RMQ_VHOST")
	RMQ_URL     string = getEnv("RMQ_URL")
)

const (
	RMQ_EXCHANGE_TYPE     string = "direct"
	RMQ_DATA_TYPE         string = "application/json"
	RMQ_CALLBACK_EXCHANGE string = "E_CALLBACK"
	RMQ_CALLBACK_QUEUE    string = "Q_CALLBACK"
)

func init() {
	// setup timezone
	loc, _ := time.LoadLocation(APP_TZ)
	time.Local = loc

	/**
	 * Listener service
	 */
	rootCmd.AddCommand(listenerCmd)

	/**
	 * Consumer service
	 */
	rootCmd.AddCommand(consumerCallbackCmd)
}

func Execute() error {
	return rootCmd.Execute()
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		log.Panicf("Error %v", key)
	}
	return value
}

// Connect to gorm mysql
func connectDb() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(URI_PGSQL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// set pool connection
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(3)
	sqlDB.SetMaxOpenConns(3)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(3 * time.Minute)

	return db, nil
}

// Connect to rabbitmq
func connectRabbitMq() (rmqp.AMQP, error) {
	var rb rmqp.AMQP
	port, _ := strconv.Atoi(RMQ_PORT)
	rb.SetAmqpURL(RMQ_HOST, port, RMQ_USER, RMQ_PASS, RMQ_VHOST)
	errConn := rb.SetUpConnectionAmqp()
	if errConn != nil {
		return rb, errConn
	}
	return rb, nil
}

// Connect to redis
func connectRedis() (*redis.Client, error) {
	opts, err := redis.ParseURL(URI_REDIS)
	if err != nil {
		return nil, err
	}
	return redis.NewClient(opts), nil
}
