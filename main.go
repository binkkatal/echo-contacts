package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/binkkatal/echo-contacts/api"
	"github.com/binkkatal/echo-contacts/dao"
)

const ()

func main() {

	db, err := dao.Connect()

	if err != nil {
		log.Fatalf("Error Connecting to database (%+v)", err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Ping Returned an error (%+v)", err)
	}
	dbApi := api.Api{DB: db}

	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	r := e.Group("/contacts")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &api.JwtCustomClaims{},
		SigningKey: []byte("a random secret key"),
	}
	r.Use(middleware.JWTWithConfig(config))
	// Login route
	e.POST("/login", api.Login)

	r.POST("/create", dbApi.Create)
	r.GET("/index", dbApi.Index)
	r.GET("/:id", dbApi.Show)
	r.PATCH("/:id/update", dbApi.Update)
	r.DELETE("/:id", dbApi.Delete)

	e.Logger.Fatal(e.Start(":8000"))
}
