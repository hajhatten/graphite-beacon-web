package main

import (
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "github.com/labstack/echo/engine/standard"
  "github.com/hajhatten/graphite-beacon-web/lib/api"
  apiTests "github.com/hajhatten/graphite-beacon-web/lib/tests/api"
)

func main() {
  e := echo.New()
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())
    
  e.File("/", "public/index.html")
  e.Static("/", "public")
  
  // Routes

  // API Alerts
  e.Get("/api/alerts", api.GetAlerts())
  e.Get("/api/alert/:id", api.GetAlert())
  e.Post("/api/alert/:id", api.NewAlert())
  e.Patch("/api/alert/:id", api.UpdateAlert())
  e.Delete("/api/alert/:id", api.DeleteAlert())
  
  // API Configurations
  e.Get("/api/configurations", api.GetConfigurations())
  e.Get("/api/configuration/:id", api.GetConfiguration())
  e.Post("/api/configuration/:id", api.NewConfiguration())
  e.Patch("/api/configuration/:id", api.UpdateConfiguration())
  e.Delete("/api/configuration/:id", api.DeleteConfiguration())
  
  // API Test
  e.Get("/api/test/mockalerts", apiTests.GetMockAlerts())
  e.Get("/api/test/configurations", apiTests.GetMockConfigurations())
  
  // API Test Create
  e.Get("/api/test/create/alerts", apiTests.CreateMockAlerts())
  e.Get("/api/test/create/configurations", apiTests.CreateMockConfigurations())

  e.Run(standard.New(":3001"))
}
