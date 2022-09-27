package netatmo

import (
	"time"

	"golang.org/x/oauth2"
)

// TokenResponse contains the authentication token received from the API
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func (t TokenResponse) Token(issueTime time.Time) *oauth2.Token {
	return &oauth2.Token{
		AccessToken:  t.AccessToken,
		TokenType:    "",
		RefreshToken: t.RefreshToken,
		Expiry:       issueTime.Add(time.Second * time.Duration(t.ExpiresIn)),
	}
}

// DeviceCollection hold all devices from netatmo account
type DeviceCollection struct {
	Body struct {
		Devices []*Device `json:"devices"`
	}
}

// Devices returns the list of devices
func (dc *DeviceCollection) Devices() []*Device {
	return dc.Body.Devices
}

// Stations is an alias of Devices
func (dc *DeviceCollection) Stations() []*Device {
	return dc.Devices()
}

// Device contains data of a station or a module.
type Device struct {
	// ID : Mac address
	ID string `json:"_id"`
	// ModuleName contains the name of the module.
	ModuleName string `json:"module_name"`
	// HomeID contains the id of the home where the station is placed.
	HomeID string `json:"home_id"`
	// HomeName contains the name of the home where the station is placed.
	HomeName string `json:"home_name"`
	// StationName contains the name of the station.
	//
	// Deprecated: Use HomeName and ModuleName instead.
	StationName string `json:"station_name"`
	// BatteryPercent : Percentage of battery remaining
	BatteryPercent *int32 `json:"battery_percent,omitempty"`
	// WifiStatus : Wifi status per Base station
	WifiStatus *int32 `json:"wifi_status,omitempty"`
	// RFStatus : Current radio status per module
	RFStatus *int32 `json:"rf_status,omitempty"`
	// Type : Module type :
	//  "NAMain" : for the base station
	//  "NAModule1" : for the outdoor module
	//  "NAModule4" : for the additional indoor module
	//  "NAModule3" : for the rain gauge module
	//  "NAModule2" : for the wind gauge module
	Type string
	// ReadOnly shows if the user owns the station.
	ReadOnly bool `json:"read_only"`
	// DashboardData : Data collection from device sensors
	DashboardData DashboardData `json:"dashboard_data"`
	// LinkedModules : Associated modules (only for station)
	LinkedModules []*Device `json:"modules"`
}

// Modules returns associated device module
func (d *Device) Modules() []*Device {
	modules := d.LinkedModules
	modules = append(modules, d)

	return modules
}

// Data returns timestamp and the list of sensor value for this module
func (d *Device) Data() (int64, map[string]interface{}) {
	// return only populate field of DashboardData
	m := make(map[string]interface{})

	if d.DashboardData.Temperature != nil {
		m["Temperature"] = *d.DashboardData.Temperature
	}
	if d.DashboardData.Humidity != nil {
		m["Humidity"] = *d.DashboardData.Humidity
	}
	if d.DashboardData.CO2 != nil {
		m["CO2"] = *d.DashboardData.CO2
	}
	if d.DashboardData.Noise != nil {
		m["Noise"] = *d.DashboardData.Noise
	}
	if d.DashboardData.Pressure != nil {
		m["Pressure"] = *d.DashboardData.Pressure
	}
	if d.DashboardData.AbsolutePressure != nil {
		m["AbsolutePressure"] = *d.DashboardData.AbsolutePressure
	}
	if d.DashboardData.Rain != nil {
		m["Rain"] = *d.DashboardData.Rain
	}
	if d.DashboardData.Rain1Hour != nil {
		m["Rain1Hour"] = *d.DashboardData.Rain1Hour
	}
	if d.DashboardData.Rain1Day != nil {
		m["Rain1Day"] = *d.DashboardData.Rain1Day
	}
	if d.DashboardData.WindAngle != nil {
		m["WindAngle"] = *d.DashboardData.WindAngle
	}
	if d.DashboardData.WindStrength != nil {
		m["WindStrength"] = *d.DashboardData.WindStrength
	}
	if d.DashboardData.GustAngle != nil {
		m["GustAngle"] = *d.DashboardData.GustAngle
	}
	if d.DashboardData.GustAngle != nil {
		m["GustAngle"] = *d.DashboardData.GustAngle
	}
	if d.DashboardData.GustStrength != nil {
		m["GustStrength"] = *d.DashboardData.GustStrength
	}

	return *d.DashboardData.LastMeasure, m
}

// Info returns timestamp and the list of info value for this module
func (d *Device) Info() (int64, map[string]interface{}) {
	// return only populate field of DashboardData
	m := make(map[string]interface{})

	// Return data from module level
	if d.BatteryPercent != nil {
		m["BatteryPercent"] = *d.BatteryPercent
	}
	if d.WifiStatus != nil {
		m["WifiStatus"] = *d.WifiStatus
	}
	if d.RFStatus != nil {
		m["RFStatus"] = *d.RFStatus
	}

	return *d.DashboardData.LastMeasure, m
}

// DashboardData is used to store sensor values
// Temperature : Last temperature measure @ LastMeasure (in °C)
// Humidity : Last humidity measured @ LastMeasure (in %)
// CO2 : Last Co2 measured @ time_utc (in ppm)
// Noise : Last noise measured @ LastMeasure (in db)
// Pressure : Last Sea level pressure measured @ LastMeasure (in mb)
// AbsolutePressure : Real measured pressure @ LastMeasure (in mb)
// Rain : Last rain measured (in mm)
// Rain1Hour : Amount of rain in last hour
// Rain1Day : Amount of rain today
// WindAngle : Current 5 min average wind direction @ LastMeasure (in °)
// WindStrength : Current 5 min average wind speed @ LastMeasure (in km/h)
// GustAngle : Direction of the last 5 min highest gust wind @ LastMeasure (in °)
// GustStrength : Speed of the last 5 min highest gust wind @ LastMeasure (in km/h)
// LastMeasure : Contains timestamp of last data received
type DashboardData struct {
	Temperature      *float32 `json:"Temperature,omitempty"` // use pointer to detect ommitted field by json mapping
	Humidity         *int32   `json:"Humidity,omitempty"`
	CO2              *int32   `json:"CO2,omitempty"`
	Noise            *int32   `json:"Noise,omitempty"`
	Pressure         *float32 `json:"Pressure,omitempty"`
	AbsolutePressure *float32 `json:"AbsolutePressure,omitempty"`
	Rain             *float32 `json:"Rain,omitempty"`
	Rain1Hour        *float32 `json:"sum_rain_1,omitempty"`
	Rain1Day         *float32 `json:"sum_rain_24,omitempty"`
	WindAngle        *int32   `json:"WindAngle,omitempty"`
	WindStrength     *int32   `json:"WindStrength,omitempty"`
	GustAngle        *int32   `json:"GustAngle,omitempty"`
	GustStrength     *int32   `json:"GustStrength,omitempty"`
	LastMeasure      *int64   `json:"time_utc"`
}
