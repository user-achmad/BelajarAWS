package main

import (
	"clean/config"
	td "clean/features/book/data"
	th "clean/features/book/handler"
	ts "clean/features/book/services"
	"clean/features/user/data"
	"clean/features/user/handler"
	"clean/features/user/services"
	"clean/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitSQL(cfg)

	userData := data.New(db)
	userService := services.NewService(userData)
	userHandler := handler.NewUserHandler(userService)

	todoData := td.New(db)
	todoService := ts.NewBookService(todoData)
	todoHandler := th.NewHandler(todoService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	routes.InitRoute(e, userHandler, todoHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
