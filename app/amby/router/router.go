package router

import (
  "fmt"
  "github.com/eifandevs/amby/controllers"
  "github.com/eifandevs/amby/interceptor"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
)

func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
  fmt.Printf("Request Body: %v\n", string(reqBody))
  fmt.Printf("Response Body: %v\n", string(resBody))
}

func Init() *echo.Echo {
	
  e := echo.New()

  e.Use(middleware.BodyDump(bodyDumpHandler))

  api := e.Group("/api")
  {
    api.GET("/accesstoken", controllers.GetAccessToken(), interceptor.BasicAuth())
    api.GET("/favorite", controllers.GetHandler())
    api.POST("/favorite", controllers.PostHandler())
  }

  e.GET("/digest", interceptor.DigestAuthenticate())

  return e
}