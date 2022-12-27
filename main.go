package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/Tomoki108/go-graphql-memo-app/graph"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	dsn := "root:pass@tcp(127.0.0.1:3306)/go_graphql_memo_app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err1 := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err1 != nil {
		log.Fatalln(err1)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", welcome())

	graphqlHandler := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{Resolvers: &graph.Resolver{DB: db}},
		),
	)
	playgroundHandler := playground.Handler("GraphQL", "/query")

	e.POST("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	err2 := e.Start(":3000")
	if err2 != nil {
		log.Fatalln(err2)
	}
}

func welcome() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome!")
	}
}
