package main

import (
	"github.com/iqbaleff214/go-fiber-api-aplikasi-pembayaran-spp/config"
	_ "github.com/iqbaleff214/go-fiber-api-aplikasi-pembayaran-spp/config"
	"github.com/iqbaleff214/go-fiber-api-aplikasi-pembayaran-spp/internal/database"
	"github.com/iqbaleff214/go-fiber-api-aplikasi-pembayaran-spp/internal/server"
	"log"
)

func main() {
	conf := config.NewConfig()

	db, err := database.NewMysqlConnection(conf.DatabaseURI())
	if err != nil {
		log.Fatalf("database connection: %v", err)
	}
	defer func() { _ = db.Close() }()

	app := server.NewServer(db)
	log.Println(app.Run(":8000"))
}
