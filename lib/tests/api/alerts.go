package tests

import (
  "log"
  "strconv"
  "net/http"
  "github.com/labstack/echo"
  "github.com/hajhatten/graphite-beacon-web/lib/types"
  "github.com/hajhatten/graphite-beacon-web/lib/db"
)

var (
  amoutOfMockAlerts = 20
)

// GetMockAlerts Returns mocked up alerts
func GetMockAlerts() echo.HandlerFunc {
  return func(c echo.Context) error {
    mockalerts := generateMockAlerts()
    return c.JSON(http.StatusOK, mockalerts)
  }
}

// CreateMockAlerts Saves mocked up alerts to db
func CreateMockAlerts() echo.HandlerFunc {
  return func(c echo.Context) error {
    mockalerts := generateMockAlerts()
    db := db.OpenDBConnection()
    
    log.Println("Checking if table alerts is present")
    if db.HasTable("alerts") == false {
      log.Println("Creating table configurations")
      db.CreateTable(&types.Alert{})
    }
    
    for _,alert := range mockalerts {
      log.Println("Saving alert " + alert.Name + " to db" )
      db.Create(&alert)
    }
    
    var alerts []types.Alert
    log.Println("Fetching alerts from db")
    db.Table("alerts").Find(&alerts)
    
    defer db.Close()
    return c.JSON(http.StatusOK, alerts)
  }
}

func generateMockAlerts() []types.Alert {
  var alerts []types.Alert
  for i := 0; i < amoutOfMockAlerts; i++ {
    alerts = append(alerts, types.Alert {
      Name: "Test Alert " + strconv.Itoa(i+1),
      Query: "testquery()",
      Source: "graphite",
      Method: "10minute",
      Interval: "10minute",
      NoData: "normal",
      Until: "",
    })
  }
  return alerts
}