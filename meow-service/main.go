package main

import (
	"fmt"
	"log"
	"meower/db"
	"time"

	"github.com/tinrab/retry"
)

type Config struct{
	PostgresDB string `envconfig:"POSTGRES_DB"`
	PostgresUser string `envconfig:"POSTGRES_USER"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD"`
	NatsAddresss string `envconfig:"NATS_ADDRESS"`
}

func main(){
	var cfg Config
	err := envconfig.Process("",&cfg)
	if err != nil{
		log.Fatal(err)
	}

	retry.ForeverSleep(2*time.Second,func(attempt int) error {
		addr := fmt.Sprintf("postgres://%s:%s@postgres/%s?sslmode=disable", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)
		repo,err := db.NewPostgres(addr)
		if err != nil{
			log.Println(err)
			return err
		}
		db.SetRepository(repo)
		return nil
	})
	defer db.CLose()
}