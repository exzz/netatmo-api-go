package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	toml "github.com/BurntSushi/toml"
	netatmo "github.com/exzz/netatmo-api-go"
)

// Command line flag
var fConfig = flag.String("f", "", "Configuration file")

// API credentials
type NetatmoConfig struct {
	ClientID     string
	ClientSecret string
	Username     string
	Password     string
}

var config NetatmoConfig

func main() {

	// Parse command line flags
	flag.Parse()
	if *fConfig == "" {
		fmt.Printf("Missing required argument -f\n")
		os.Exit(0)
	}

	if _, err := toml.DecodeFile(*fConfig, &config); err != nil {
		fmt.Printf("Cannot parse config file: %s\n", err)
		os.Exit(1)
	}

	n, err := netatmo.NewClient(netatmo.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Username:     config.Username,
		Password:     config.Password,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dc, err := n.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ct := time.Now().UTC().Unix()

	for _, station := range dc.Stations() {
		fmt.Printf("Station : %s\n", station.StationName)
		fmt.Printf("\tCity: %s\n\tCountry: %s\n\tTimezone: %s\n\tLongitude: %f\n\tLatitude: %f\n\tAltitude: %d\n\n", station.Place.City, station.Place.Country, station.Place.Timezone, station.Place.Location.Longitude, station.Place.Location.Latitude, station.Place.Altitude)

		for _, module := range station.Modules() {
			fmt.Printf("\tModule : %s\n", module.ModuleName)

			{
				if module.DashboardData.LastMeasure == nil {
					fmt.Printf("\t\tSkipping %s, no measurement data available.\n", module.ModuleName)
					continue
				}
				ts, data := module.Info()
				for dataName, value := range data {
					fmt.Printf("\t\t%s : %v (updated %ds ago)\n", dataName, value, ct-ts)
				}
			}

			{
				ts, data := module.Data()
				for dataName, value := range data {
					fmt.Printf("\t\t%s : %v (updated %ds ago)\n", dataName, value, ct-ts)
				}
			}
		}
	}
}
