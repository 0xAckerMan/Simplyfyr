package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

)

var Version = "1.0.0"

type Config struct{
    env string
    port int
    db struct{
        dsn string
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
    flag.Parse()

    logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

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
    err := srv.ListenAndServe()
    if err != nil {
        logger.Fatal(err)
    }
}
