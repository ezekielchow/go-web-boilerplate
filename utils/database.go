package utils

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var pool *sql.DB

func SetupDatabase(dsn string) {

	if len(dsn) == 0 {
		log.Fatal("Datasource is not set")
		panic("Datasource is not set")
	}

	var err error
	pool, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("unable to use data source name", err)
		panic("Unable to load database")
	}
	defer pool.Close()

	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(3)
	pool.SetMaxOpenConns(3)

	ctx, stop := context.WithCancel(context.Background())

	defer stop()

	appSignal := make(chan os.Signal, 3)
	signal.Notify(appSignal, os.Interrupt)

	go func() {
		<-appSignal
		stop()
	}()

	Ping(ctx)

}

func Ping(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err := pool.PingContext(ctx); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
		panic("unable to connect to database")
	}
}
