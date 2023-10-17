package main

import (
	"log"
	"os"

	"antin0.de/comm-relay/handlers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mysqlDsn := os.Getenv("MYSQL_DSN")
	db, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	h := handlers.HandlerParams{Db: db}

	r := gin.Default()
	cookieSecret := os.Getenv("COOKIE_SECRET")
	store := cookie.NewStore([]byte(cookieSecret))
	r.Use(sessions.Sessions("session", store))

	r.GET("/ping", h.Ping())

	listenAddress := os.Getenv("LISTEN_ADDRESS")
	r.Run(listenAddress)
}
