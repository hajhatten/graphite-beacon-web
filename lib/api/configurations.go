package api

import (
  "log"
  "strconv"
  "net/http"
  "github.com/labstack/echo"
  "github.com/hajhatten/graphite-beacon-web/lib/db"
  "github.com/hajhatten/graphite-beacon-web/lib/types"
)

// GetConfigurations Returns all the saved configurations
func GetConfigurations() echo.HandlerFunc {
  return func(c echo.Context) error {
    db := db.OpenDBConnection()
    var configurations []types.Configuration
    
    if db.HasTable("configurations") == false {  
      log.Println("table alerts not found, creating it")
      db.CreateTable(&configurations)
    }
    
    log.Println("fetching configurations from db")
    db.Find(&configurations)
    
    defer db.Close()
    return c.JSON(http.StatusOK, configurations)
  }
}

// GetDeletedConfigurations Returns all the saved alerts
func GetDeletedConfigurations() echo.HandlerFunc {
  return func(c echo.Context) error {
    db := db.OpenDBConnection()
    var configurations []types.Configuration

    if db.HasTable("configurations") == false {  
      log.Println("table configurations not found, creating it")
      db.CreateTable(&configurations)
    }
    
    log.Println("fetching configurations from db")
    db.Unscoped().Find(&configurations)
    
    defer db.Close()
    return c.JSON(http.StatusOK, configurations)
  }
}

// GetConfiguration Returns the configuration by id
func GetConfiguration() echo.HandlerFunc {
  return func(c echo.Context) error {
    db := db.OpenDBConnection()
    var configuration types.Configuration
    
    log.Println("fetching configuration from db")
    db.First(&configuration, c.Param("id"))
    
    defer db.Close()
    return c.JSON(http.StatusOK, configuration)
  }
}

// NewConfiguration Creates new configuration
func NewConfiguration() echo.HandlerFunc {
  return func(c echo.Context) error {
    db := db.OpenDBConnection()
    var configuration types.Configuration
    
    log.Println("updating configuration object from form")
    configuration.Name                = c.FormValue("name")
    configuration.GraphiteURL         = c.FormValue("graphite_url")
    configuration.PublicGraphiteURL   = c.FormValue("public_graphite_url")
    configuration.AuthUsername        = c.FormValue("auth_username")
    configuration.AuthPassword        = c.FormValue("auth_password")
    configuration.Pidfile             = c.FormValue("pidfile")
    configuration.Format              = c.FormValue("format")
    configuration.Interval            = c.FormValue("interval")
    configuration.TimeWindow          = c.FormValue("time_window")
    configuration.RepeatInterval      = c.FormValue("repeat_interval")
    configuration.Until               = c.FormValue("until")
    configuration.Logging             = c.FormValue("logging")
    configuration.Method              = c.FormValue("method")
    configuration.NoData              = c.FormValue("no_data")
    configuration.LoadingError        = c.FormValue("loading_error")
    configuration.Prefix              = c.FormValue("prefix")
    configuration.SendInitial, _      = strconv.ParseBool(c.FormValue("send_initial"))
    configuration.DefaultNanValue, _  = strconv.ParseInt(c.FormValue("default_nan_value"), 10, 32)
    configuration.IgnoreNan, _        = strconv.ParseBool(c.FormValue("ignore_nan"))
    
    log.Println("saving new configuration to db")
    db.Create(&configuration)
    
    defer db.Close()
    return c.JSON(http.StatusOK, configuration)
  }
}

// UpdateConfiguration Updates the configuration by id
func UpdateConfiguration() echo.HandlerFunc {
  return func(c echo.Context) error {
    db := db.OpenDBConnection()
    var configuration types.Configuration
    
    log.Println("fetching configuration from db")
    db.First(&configuration, c.Param("id"))
    
    log.Println("updating configuration object from form")
    configuration.Name = c.FormValue("name")
    configuration.GraphiteURL = c.FormValue("graphite_url")
    configuration.PublicGraphiteURL = c.FormValue("public_graphite_url")
    configuration.AuthUsername = c.FormValue("auth_username")
    configuration.AuthPassword = c.FormValue("auth_password")
    configuration.Pidfile = c.FormValue("pidfile")
    configuration.Format = c.FormValue("format")
    configuration.Interval = c.FormValue("interval")
    configuration.TimeWindow = c.FormValue("time_window")
    configuration.RepeatInterval = c.FormValue("repeat_interval")
    configuration.Until = c.FormValue("until")
    configuration.Logging = c.FormValue("logging")
    configuration.Method = c.FormValue("method")
    configuration.NoData = c.FormValue("no_data")
    configuration.LoadingError = c.FormValue("loading_error")
    configuration.Prefix = c.FormValue("prefix")
    configuration.SendInitial, _ = strconv.ParseBool(c.FormValue("send_initial"))
    configuration.DefaultNanValue, _ = strconv.ParseInt(c.FormValue("default_nan_value"), 10, 32)
    configuration.IgnoreNan, _ = strconv.ParseBool(c.FormValue("ignore_nan"))
    
    log.Println("saving updated configuration to db")
    db.Save(&configuration)
    
    defer db.Close()
    return c.Render(http.StatusOK, "message", "configuration '" + configuration.Name + "' updated.")
  }
}

// DeleteConfiguration Deletes the configuration by id
func DeleteConfiguration() echo.HandlerFunc {
  return func(c echo.Context) error {
    db := db.OpenDBConnection()
    var configuration types.Configuration
    
    log.Println("deleting configuration from db")
    db.Delete(&configuration, c.Param("id"))
    
    defer db.Close()
    return c.NoContent(204)
  }
}

// UndeleteConfiguration undeletes an alert
func UndeleteConfiguration() echo.HandlerFunc {
  return func(c echo.Context) error {
    db := db.OpenDBConnection()
    var configuration types.Configuration
    
    configuration.DeletedAt = nil
    
    defer db.Close()
    return c.NoContent(204)
  }
}

// PermDeleteConfiguration Deletes the alert by id
func PermDeleteConfiguration() echo.HandlerFunc {
  return func(c echo.Context) error {
    db := db.OpenDBConnection()
    var configuration types.Configuration
    
    log.Println("fetching configuration from db")
    db.First(&configuration, c.Param("id"))
    
    log.Println("permanently deleting configuration from db")
    db.Unscoped().Delete(&configuration)
    
    defer db.Close()
    return c.NoContent(204)
  }
}