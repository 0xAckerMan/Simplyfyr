package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var Version = "1.0.0"

type Config struct{
    env string
    port int
    db struct{
        dsn string
        maxOpenConns int
        maxIdleConns int
        maxIdleTime string
    }
}

type Application struct{
    config Config
    logger *log.Logger
}

func main(){
    var cfg Config

    flag.IntVar(&cfg.port, "port", 4000, "API running port")
    flag.StringVar(&cfg.env, "env", "dev", "The running environment (dev|stag|prod)")
    flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("SIMPLIFYR_DB_DSN"), "Postgres DSN string")
    flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "postgres maximum idle time")
    flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "postgres maximum number of open connections")
    flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "postgres maximum idle time")
    flag.Parse()

    logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

    db, err := openDB(cfg)
    if err != nil{
        logger.Fatal(err)
    }

    
    defer db.Close()

    logger.Println("database connection pool created established")

    app := &Application{
        config: cfg,
        logger: logger,
    }

    addr := fmt.Sprintf(":%d", cfg.port)

    srv := http.Server{
        Addr: addr,
        Handler: app.routes(),
        ReadTimeout: 30 * time.Second,
        WriteTimeout: 30 * time.Second,
        IdleTimeout: time.Minute,
    }
    logger.Printf("Starting %s server on port %d", cfg.env, cfg.port)
    err = srv.ListenAndServe()
    if err != nil {
        logger.Fatal(err)
    }
}

func openDB(cfg Config) (*sql.DB, error){
    db, err := sql.Open("postgres", cfg.db.dsn)
    if err != nil {
        return nil, err
    }

    db.SetMaxOpenConns(cfg.db.maxOpenConns)
    db.SetMaxIdleConns(cfg.db.maxIdleConns)

    duration, err := time.ParseDuration(cfg.db.maxIdleTime)
    if err!= nil{
        return nil, err
    }

    db.SetConnMaxIdleTime(duration)
    
    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()
    err = db.PingContext(ctx)
    if err != nil {
        return nil, err
    }


    return db, nil
}
