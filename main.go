package main

import (
  "github.com/rs/cors"
  "github.com/labstack/echo"
  mw "github.com/labstack/echo/middleware"
  "github.com/labstack/echo/engine/standard"
  "github.com/hajhatten/graphite-beacon-web/lib/api"
  apiTests "github.com/hajhatten/graphite-beacon-web/lib/tests/api"
)

func main() {
  e := echo.New()
  e.Use(mw.Logger())
  e.Use(mw.Recover())
  cors := standard.WrapMiddleware(cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	}).Handler)
  e.Use(cors)


  // Static files
  e.File("/", "public/index.html")
  e.Static("/", "public")

  // Routes

  // API Alerts
  alertsAPI := e.Group("/api/alert")
  alertsAPI.Get("s", api.GetAlerts())
  alertsAPI.Get("s/deleted", api.GetDeletedAlerts())
  alertsAPI.Get("/:id", api.GetAlert())
  alertsAPI.Post("/:id", api.NewAlert())
  alertsAPI.Patch("/:id", api.UpdateAlert())
  alertsAPI.Delete("/:id", api.DeleteAlert())
  alertsAPI.Put("/:id/undelete", api.UndeleteAlert())
  alertsAPI.Delete("/:id/permdelete", api.PermDeleteAlert())

  alertsAPI.Use(cors)

  // API Configurations
  configurationssAPI := e.Group("/api/configurations")
  configurationssAPI.Get("", api.GetConfigurations())

  configurationssAPI.Get("/:id", api.GetConfiguration())
  configurationssAPI.Post("/:id", api.NewConfiguration())
  configurationssAPI.Put("/:id", api.UpdateConfiguration())
  configurationssAPI.Delete("/:id", api.DeleteConfiguration())
  configurationssAPI.Patch("/:id/undelete", api.UndeleteConfiguration())
  configurationssAPI.Delete("/:id/permdelete", api.PermDeleteConfiguration())

  configurationssAPI.Use(cors)

  // API Test
  e.Get("/api/test/mockalerts", apiTests.GetMockAlerts())
  e.Get("/api/test/configurations", apiTests.GetMockConfigurations())

  // API Test Create
  e.Get("/api/test/create/alerts", apiTests.CreateMockAlerts())
  e.Get("/api/test/create/configurations", apiTests.CreateMockConfigurations())

  // CORS
  e.Use(cors)

  e.Run(standard.New(":3001"))
}
