package main

import (
	"flag"
	"net/http"

	"github.com/arunvm/recro/config"
	"github.com/arunvm/recro/models"
	"github.com/rs/cors"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

type server struct {
	db     *gorm.DB
	routes *gin.Engine
}

func newServer() *server {
	s := server{}
	return &s
}

func main() {
	server := newServer()

	// Logging options
	log.SetFormatter(&log.JSONFormatter{})

	// Reading file path from flag
	filePath := flag.String("config-path", "config.yaml", "filepath to configuration file")
	flag.Parse()

	// Reading config variables
	config, err := config.Initialise(*filePath)
	if err != nil {
		log.Fatalf("Failed to read config\n%v", err)
	}

	// "host=myhost port=myport user=gorm dbname=gorm password=mypassword"
	server.db, err = gorm.Open("postgres", "host= "+config.Database.Host+" port="+config.Database.Port+" user="+config.Database.User+" dbname="+config.Database.DatabaseName+" password="+config.Database.Password+" sslmode=disable")
	if err != nil {
		panic(err)
	}

	// server.db.LogMode(true)
	models.MigrateDB(server.db)

	// Setting up routes
	server.routes = initialiseRoutes(server)
	routes := cors.AllowAll().Handler(server.routes)

	http.ListenAndServe(":5000", routes)
}
