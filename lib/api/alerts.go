package api

import (
  "log"
  "net/http"
  "github.com/labstack/echo"
  "github.com/hajhatten/graphite-beacon-web/lib/db"
  "github.com/hajhatten/graphite-beacon-web/lib/types"
)


// GetAlerts Returns all the saved alerts
func GetAlerts() echo.HandlerFunc {
  return func(c echo.Context) error {
    db := db.OpenDBConnection()
    var alerts []types.Alert

    if db.HasTable("alerts") == false {  
      log.Println("table alerts not found, creating it")
      db.CreateTable(&alerts)
    }
    
    log.Println("fetching alerts from db")
    db.Find(&alerts)
    
    defer db.Close()
    return c.JSON(http.StatusOK, alerts)
  }
}

// GetDeletedAlerts Returns all the saved alerts
func GetDeletedAlerts() echo.HandlerFunc {
  return func(c echo.Context) error {
    db := db.OpenDBConnection()
    var alerts []types.Alert

    if db.HasTable("alerts") == false {  
      log.Println("table alerts not found, creating it")
      db.CreateTable(&alerts)
    }
    
    log.Println("fetching alerts from db")
    db.Unscoped().Where("deleted_at IS NOT ?", nil).Find(&alerts)
    
    defer db.Close()
    return c.JSON(http.StatusOK, alerts)
  }
}

// GetAlert Returns the alert by id
func GetAlert() echo.HandlerFunc {
  return func(c echo.Context) error {
    db := db.OpenDBConnection()
    var alert types.Alert
    
    log.Println("fetching alert from db")
    db.First(&alert, c.Param("id"))
    
    defer db.Close()
    return c.JSON(http.StatusOK, alert)
  }
}

// NewAlert Creates new alert
func NewAlert() echo.HandlerFunc {
  return func(c echo.Context) error {
    db := db.OpenDBConnection()
    var alert types.Alert
    
    log.Println("updating alert object from form")
    alert.Name = c.FormValue("name")
    alert.Query = c.FormValue("query")
    alert.Source = c.FormValue("source")
    alert.Format = c.FormValue("format")
    alert.Method = c.FormValue("method")
    alert.Interval = c.FormValue("interval")
    alert.NoData = c.FormValue("no_data")
    alert.Until = c.FormValue("until")
    
    log.Println("saving new alert to db")
    db.Create(&alert)
    
    defer db.Close()
    return c.JSON(http.StatusOK, alert)
  }
}

// UpdateAlert Updates the alert by id
func UpdateAlert() echo.HandlerFunc {
  return func(c echo.Context) error {
    db := db.OpenDBConnection()
    var alert types.Alert
    
    log.Println("fetching alert from db")
    
    log.Println("updating alert object from form")
    alert.Name = c.FormValue("name")
    alert.Query = c.FormValue("query")
    alert.Source = c.FormValue("source")
    alert.Format = c.FormValue("format")
    alert.Method = c.FormValue("method")
    alert.Interval = c.FormValue("interval")
    alert.NoData = c.FormValue("no_data")
    alert.Until = c.FormValue("until")
    
    log.Println("saving updated alert to db")
    db.Save(&alert)
    
    defer db.Close()
    return c.Render(http.StatusOK, "message", "alert '" + alert.Name + "' updated.")
  }
}

// DeleteAlert Deletes the alert by id
func DeleteAlert() echo.HandlerFunc {
  return func(c echo.Context) error {
    db := db.OpenDBConnection()
    var alert types.Alert
    
    log.Println("fetching alert from db")
    db.First(&alert, c.Param("id"))
    
    log.Println("deleting alert from db")
    db.Delete(&alert)
    
    defer db.Close()
    return c.NoContent(204)
  }
}

// UndeleteAlert undeletes an alert
func UndeleteAlert() echo.HandlerFunc {
  return func(c echo.Context) error {
    db := db.OpenDBConnection()
    var alert types.Alert
    
    log.Println("setting alert deleted_at to nil in db")
    db.Unscoped().First(&alert, c.Param("id")).Update("deleted_at", nil)
    
    defer db.Close()
    return c.NoContent(204)
  }
}

// PermDeleteAlert Deletes the alert by id
func PermDeleteAlert() echo.HandlerFunc {
  return func(c echo.Context) error {
    db := db.OpenDBConnection()
    var alert types.Alert
    
    log.Println("fetching alert from db")
    db.First(&alert, c.Param("id"))
    
    log.Println("deleting alert from db")
    db.Unscoped().Delete(&alert)
    
    defer db.Close()
    return c.NoContent(204)
  }
}