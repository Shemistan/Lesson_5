package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	var err error
	fmt.Println("До----> ", os.Getenv("TEST_HOST"))

	_ = os.Setenv("TEST_HOST", "localhost:999999")

	fmt.Println("TEST_HOST----> ", os.Getenv("TEST_HOST"))

	// Экспорт env-файла в переменные окружения
	err = godotenv.Load("./dev/local.env")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("POSTGRES_HOST----> ", os.Getenv("POSTGRES_HOST"))

	// Выкачиваем переменные окружения в структуру
	var cfg EnvConfig
	err = envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v ", cfg)

}

type EnvConfig struct {
	Host string `envconfig:"HTTP_PORT" default:"localhost"`
	Port int    `envconfig:"HTTP_HOST" default:"8080"`

	Postgres Postgres `envconfig:"POSTGRES"`
}

type Postgres struct {
	User            string        `envconfig:"USER" required:"true"`
	Password        string        `envconfig:"PASSWORD" required:"true"`
	Host            string        `envconfig:"HOST" required:"true"`
	Database        string        `envconfig:"DATABASE" required:"true"`
	MaxIdleConnTime time.Duration `envconfig:"MAX_IDLE_CONN_TIME" default:"5m"`
	MaxConns        int           `envconfig:"MAX_CONNS" default:"20"`
	ConnMaxLifetime time.Duration `envconfig:"CONN_MAX_LIFETIME" default:"10m"`
}
