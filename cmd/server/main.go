package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"template-api-pg/internal/api"
	_ "template-api-pg/internal/config"
	"template-api-pg/internal/storage"
	"time"

	"github.com/spf13/viper"
)

func main() {
	log.Println("server starting")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	conn, err := storage.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	server := api.NewServer(ctx, conn)

	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		log.Println("shutting down server")
		server.Shutdown(ctx)
	}()

	log.Println("db connected")
	log.Printf("api server started on :%d", viper.GetInt("api_port"))
	server.ListenAndServe()
	log.Println("server shut down")
}
