package types

import (
  "net/url"
  "github.com/jinzhu/gorm"
  _ "encoding/json" //enable json output tagging
)

// Configuration define configuration and defaults for alarms
type Configuration struct {
  gorm.Model
  Name              string    `gorm:"index;not null;unique" json:"name"` // Configuration Name
  GraphiteURL       string    `json:"graphite_url"`         // Graphite server URL
  PublicGraphiteURL string    `json:"public_graphite_url"`  // Public graphite server URL. Used when notifying handlers, defaults to graphite_url
  AuthUsername      string    `json:"auth_username"`        // HTTP AUTH username
  AuthPassword      string    `json:"auth_password"`        // HTTP AUTH password
  Pidfile           string    `json:"pidfile"`              // Path to a pidfile
  Format            string    `json:"format"`               // Default values format (none, bytes, s, ms, short) Can be redefined for each alert.
  Interval          string    `json:"interval"`             // Default query interval. Can be redefined for each alert.
  TimeWindow        string    `json:"time_window"`          // Default time window for Graphite queries. Defaults to query interval, can be redefined for each alert.
  RepeatInterval    string    `json:"repeat_interval"`      // Notification repeat interval. If an alert is failed, its notification will be repeated with the interval below.
  Until             string    `json:"until"`                // Default end time for Graphite queries. Defaults to the current time, can be redefined for each alert.
  Logging           string    `json:"logging"`              // Default log level
  Method            string    `json:"method"`               // Default method (average, last_value, sum, minimum, maximum). Can be redefined for each alert.
  NoData            string    `json:"no_data"`              // Default alert to send when no data received (normal = no alert). Can be redefined for each alert.
  LoadingError      string    `json:"loading_error"`        // Default alert to send when loading failed (timeout, server error, etc). normal = no alert. Can be redefined for each alert.
  Prefix            string    `json:"prefix"`               // Default prefix (used for notifications)
  SendInitial       bool      `gorm:"default:false" json:"send_initial"`  // Send initial values (Send current values when reactor starts)
  DefaultNanValue   int64    `json:"default_nan_value"`    // Used together to ignore the missing value
  IgnoreNan         bool      `json:"ignore_nan"`           // Used together to ignore the missing value
}


// Alert defines specific query and querying parameters
type Alert struct {
  gorm.Model
  Name            string      `gorm:"index;not null;unique" json:"name"`  // Name of alert
  Query           string      `json:"query"`      // Default alert to send when no data received (normal = no alert). Can be defined as default in Configuration.
  Source          string      `json:"source"`     // Source name
  Format          string      `json:"format"`     // Values format (none, bytes, s, ms, short) Can be redefined for each alert.
  Method          string      `json:"method"`     // Method (average, last_value, sum, minimum, maximum). Can be redefined for each alert.
  Interval        string      `json:"interval"`   // Query interval. Can be defined as default in Configuration.
  NoData          string      `json:"no_data"`    // Alert to send when no data received (normal = no alert). Can be defined as default in Configuration.
  Until           string      `json:"until"`      // End time for Graphite queries. Defaults to the current time. Can be defined as default in Configuration.
}

// Rule defines specific level, operator and value to trigger on 
type Rule struct {
  gorm.Model
  AlertID   uint                      // Rule belongs to alert with this ID 
  Level     string `json:"level"`     // warning, critical
  Operator  string `json:"operator"`  // operator like lessthan, morethan and equals (<, >, ==)
  Value     string `json:"value"`
}

// Handlers

// EmailHandler defines how to send email notifications
type EmailHandler struct {
  AlertID     uint      // Handler belongs to configuration with this ID
  From        string    `json:"from"`       // From address
  To          []string  `json:"to"`         // List of email addresses to send to
  Host        string    `json:"host"`       // SMTP host
  Port        int       `json:"port"`       // SMTP port
  Username    string    `json:"username"`   // SMTP user (optional)
  Password    string    `json:"password"`   // SMTP password (optional)
  UseTLS      bool      `json:"use_tls"`    // Use TLS?
  HTML        bool      `json:"html"`       // Send HTML emails?
}

// HipChatHandler defines how to send hipchat notifications
type HipChatHandler struct {
  AlertID     uint      // Handler belongs to configuration with this ID
  URL   url.URL         `json:"url"`
  Room  string          `json:"room"`
  Key   string          `json:"key"`
}

// WebhooktHandler defines how to send webhook notifications
type WebhooktHandler struct {
  AlertID     uint      // Handler belongs to configuration with this ID
  URL     url.URL       `json:"url"`
  Params  interface{}   `json:"params"`
  Method  string        `json:"method"`
}

// SlackHandler defines how to send slack notifications
type SlackHandler struct {
  AlertID     uint      // Handler belongs to configuration with this ID
  Webhook   url.URL     `json:"webhook"`
  Channel   string      `json:"channel"`
  Username  string      `json:"username"`
}

// CLIHandler defines how to send cli commands when triggered by alert
type CLIHandler struct {
  AlertID     uint      // Handler belongs to configuration with this ID
  Command         string    `json:"command"`
  AlertsWhitelist []string  `json:"alert_whitelist"`
}

// PagerDutyHandler defines how to send pagerduty notifications
type PagerDutyHandler struct {
  AlertID     uint      // Handler belongs to configuration with this ID
  Subdomain   string  `json:"subdomain"`
  Apitoken    string  `json:"apitoken"`
  ServiceKey  string  `json:"service_key"`
}

// TelegramHandler defines how to send telegram notifications
type TelegramHandler struct {
  AlertID     uint      // Handler belongs to configuration with this ID
  Token       string  `json:"token"`
  BotIdent    string  `json:"bot_ident"`
}
