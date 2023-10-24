package main

import (
	"database/sql"
	"log"
	httpHandler "project_modul_name/handler"

	_ "github.com/go-sql-driver/mysql" //import for side effect
	"github.com/labstack/echo/v4"
)

func main() {
	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "root"
	dbPass := ""
	dbName := "article"

	dsn := dbUser + `:` + dbPass + `@tcp(` + dbHost + `:` + dbPort + `)/` + dbName + `?parseTime=1&loc=Asia%2FJakarta`

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	handler := httpHandler.InitArticle(db)
	echoServer := echo.New()

	//Register the handler
	echoServer.GET("/articles", handler.FetchArticles)

	//Start the server
	echoServer.Start(":9090")
}
