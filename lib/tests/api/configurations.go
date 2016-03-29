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
  amoutOfMockConfs = 20
)

// GetMockConfigurations Returns mocked up configurations
func GetMockConfigurations() echo.HandlerFunc {
  return func(c echo.Context) error {
    mockconfigurations := generateMockConfigurations()
    return c.JSON(http.StatusOK, mockconfigurations)
  }
}

// CreateMockConfigurations Saves mocked up configurations to db
func CreateMockConfigurations() echo.HandlerFunc {
  return func(c echo.Context) error {
    mockconfigurations := generateMockConfigurations()
    db := db.OpenDBConnection()
    
    log.Println("Checking if table configurations is present")
    if db.HasTable("configurations") == false {
      log.Println("Table configurations is not present")
      log.Println("Creating table configurations")
      db.CreateTable(&types.Configuration{})
    } else {
      log.Println("Table configurations is present")
    }
    
    for _,configuration := range mockconfigurations {
      log.Println("Saving configuration " + configuration.Name + " to db" )
      db.Create(&configuration)
    }
    
    var configurations []types.Configuration
    log.Println("Fetching configurations from db")
    db.Table("configurations").Find(&configurations)
    
    defer db.Close()
    return c.JSON(http.StatusOK, configurations)
  }
}

func generateMockConfigurations() []types.Configuration {
  var configurations []types.Configuration  
  for i := 0; i < amoutOfMockConfs; i++ {
    configurations = append(configurations, types.Configuration {
      Name: "Test Configuration " + strconv.Itoa(i+1),
      GraphiteURL: "http://localhost:81",
      PublicGraphiteURL: "http://localhost:80",
      AuthUsername: "graphite",
      AuthPassword: "10minute",
      Pidfile: "/var/run/graphitebeacon",
      Format: "short",
      Interval: "10minute",
      TimeWindow: "10minute",
      RepeatInterval: "10minute",
      Until: "1000",
      Logging: "info",
      Method: "average",
      NoData: "normal",
      Prefix: "prefix_",
      SendInitial: false,
      DefaultNanValue: 1,
    })
  }
  return configurations
}