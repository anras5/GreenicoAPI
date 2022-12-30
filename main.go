package main

import (
	"database/sql"
	"fmt"
	"github.com/anras5/GreenicoAPI/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

func main() {

	err := godotenv.Load()
	host := os.Getenv("HOSTNAME")
	dbname := os.Getenv("DBNAME")
	user := os.Getenv("DBUSER")
	port := os.Getenv("PORT")
	password := os.Getenv("DB_PASSWD")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// connect to database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.LoadHTMLGlob("templates/*/**")

	// API
	router.GET("/api/reading", handlers.GetReading(db))
	router.GET("/api/readings", handlers.GetReadings(db))
	router.GET("/api/all-readings", handlers.GetAllReadings(db))
	router.POST("/api/reading", handlers.InsertReading(db))
	router.DELETE("/api/reading", handlers.DeleteReading(db))

	// WEB APP
	router.GET("/", handlers.MainRender(db))
	router.GET("/voc", handlers.PageRender(db, "voc"))
	router.GET("/temperature", handlers.PageRender(db, "temperature"))
	router.GET("/humidity", handlers.PageRender(db, "humidity"))
	router.GET("/light", handlers.PageRender(db, "light"))
	router.GET("/pressure", handlers.PageRender(db, "pressure"))
	router.GET("/uv", handlers.PageRender(db, "uv"))
	router.Static("/static", "./static/")
	err = router.Run(":8080")
	if err != nil {
		return
	}
}
