package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"server/models"
	"server/utils/env"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const version = "1.0"

type config struct {
	port string
	env  string
	db   struct {
		dsn string
	}
}

type appstatus struct {
	Status      string
	Environment string
	Version     string
}

type application struct {
	config config
	logger *log.Logger
	models models.Models
}

func main() {
	var cfg config

	env.CheckENV()

	flag.StringVar(&cfg.port, "port", env.MustGet("PORT"), "Server port will be listen on")
	flag.StringVar(&cfg.env, "environment", env.MustGet("ENVIRONMENT"), "Application environment")
	flag.StringVar(&cfg.db.dsn, "dsn", env.MustGet("MONGODB_URI"), "Database URI")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	//Database connection
	db, err := openDB(cfg)

	if err != nil {
		logger.Fatal("Error Connecting to db")
	} else {
		logger.Println("Connected to database")
	}

	defer db.Disconnect(context.TODO())

	app := &application{
		config: cfg,
		logger: logger,
		models: models.NewModels(db),
	}

	logger.Println("Pepare to start")

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%v", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
	}

	logger.Println("Server is running on port ", cfg.port)

	err = srv.ListenAndServe()

	if err != nil {
		logger.Println(err)
	}

}

func openDB(cfg config) (*mongo.Client, error) {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.db.dsn))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	return client, nil
}
