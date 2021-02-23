package main

import (
	"context"
	"fmt"
	"github.com/gitalek/taxi/calc/config"
	"github.com/gitalek/taxi/calc/server"
	//"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	//"github.com/pressly/goose"
	"github.com/spf13/viper"
	"log"
	"net/url"
)

type DbConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	Timeout  int
}

func NewPoolConfig(cfg *DbConfig) (*pgxpool.Config, error) {
	connStr := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=disable&connect_timeout=%d",
		"postgres",
		url.QueryEscape(cfg.Username),
		url.QueryEscape(cfg.Password),
		cfg.Host,
		cfg.Port,
		cfg.DbName,
		cfg.Timeout,
	)

	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}
	return poolConfig, nil
}

func NewConnection(poolConfig *pgxpool.Config) (*pgxpool.Pool, error) {
	conn, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}
	return conn, err
}

func main() {
	// reading config data into viper
	err := config.Init()
	if err != nil {
		log.Fatalf("error while reading config: %#v\n", err)
	}

	cfg := &DbConfig{
		Host:     "db",
		Port:     "5432",
		Username: "user",
		Password: "password",
		DbName:   "taxi_db",
		Timeout:  5,
	}
	poolConfig, err := NewPoolConfig(cfg)
	if err != nil {
		log.Fatalf("Pool config error: %#v\n", err)
	}
	poolConfig.MaxConns = 5
	c, err := NewConnection(poolConfig)
	if err != nil {
		log.Fatalf("Connect to database failed: %#v\n", err)
	}
	log.Printf("Connection OK!\n")

	_, err = c.Exec(context.Background(), ";")
	if err != nil {
		log.Fatalf("Ping failed: %#v\n", err)
	}
	fmt.Printf("Ping OK!\n")

	rate_name := "business"
	var taxiServicePrice, minPrice, minuteRate, meterRate float64
	err = c.QueryRow(context.Background(), "SELECT taxi_service, min_price, minute_rate, meter_rate FROM rates WHERE name=$1 LIMIT 1;", rate_name).
		Scan(&taxiServicePrice, &minPrice, &minuteRate, &meterRate)
	if err != nil {
		log.Fatalf("error while selecting from rates table: %#v\n", err)
	}
	log.Printf(
		"taxiServicePrice -> %#v, minPrice -> %#v, minuteRate -> %#v, meterRate -> %#v\n",
		taxiServicePrice, minPrice, minuteRate, meterRate,
	)

	// names collision
	//config := server.AppConfig{
	//	//todo приведение типа?
	//	//todo проверить на пустые поля
	//	Port:             viper.GetString("port"),
	//	ApiUrl:           viper.GetString("apiUrl"),
	//	TaxiServicePrice: viper.GetFloat64("taxiServicePrice"),
	//	MinPrice:         viper.GetFloat64("minPrice"),
	//	MinuteRate:       viper.GetFloat64("minuteRate"),
	//	MeterRate:        viper.GetFloat64("meterRate"),
	//}
	config := server.AppConfig{
		//todo приведение типа?
		//todo проверить на пустые поля
		Port:             viper.GetString("port"),
		ApiUrl:           viper.GetString("apiUrl"),
		TaxiServicePrice: taxiServicePrice,
		MinPrice:         minPrice,
		MinuteRate:       minuteRate,
		MeterRate:        meterRate,
		DB:               c,
	}

	app := server.NewApp(config)
	if err := app.Run(); err != nil {
		log.Fatalf("error while running server: %#v", err.Error())
	}
}
