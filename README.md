# netatmo-api-go
Simple API to access Netatmo weather station data written in Go.

Currently tested only with one weather station, outdoor and indoor modules and rain gauge. Let me know if it works with wind gauge.

## Quickstart

- [Create a new netatmo app](https://dev.netatmo.com/dev/createapp)
- Download module ```go get github.com/exzz/netatmo-api-go```
- Try below example (do not forget to edit auth credentials)

## Example

```
go
package main

import (
	"fmt"
	"os"
	"time"

	netatmo "github.com/exzz/netatmo-api-go"
)

func main() {

	n, err := netatmo.NewClient(netatmo.Config{
    ClientID:     "YOUR_APP_ID",
    ClientSecret: "YOUR_APP_SECRET",
    Username:     "YOUR_CREDENTIAL",
    Password:     "YOUR_PASSWORD",
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

		for _, module := range station.Modules() {
			fmt.Printf("\tModule : %s\n", module.ModuleName)

			{
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
```

Output should look like this :
```
Station : Home               
        Module : Chambre Elsa                              
                BatteryPercent : 47 (updated 323s ago)     
                RFStatus : 68 (updated 323s ago)           
                Temperature : 22.8 (updated 323s ago)      
                Humidity : 53 (updated 323s ago)           
                CO2 : 446 (updated 323s ago)               
        Module : Chambre parents                           
                BatteryPercent : 50 (updated 323s ago)     
                RFStatus : 71 (updated 323s ago)           
                Temperature : 19.9 (updated 323s ago)      
                Humidity : 61 (updated 323s ago)           
                CO2 : 428 (updated 323s ago)               
        Module : Chambre Jules                             
                BatteryPercent : 46 (updated 323s ago)     
                RFStatus : 60 (updated 323s ago)           
                CO2 : 396 (updated 323s ago)               
                Temperature : 22 (updated 323s ago)        
                Humidity : 54 (updated 323s ago)           
        Module : Exterieur   
                BatteryPercent : 37 (updated 323s ago)     
                RFStatus : 66 (updated 323s ago)           
                Temperature : 23.4 (updated 323s ago)      
                Humidity : 52 (updated 323s ago)           
        Module : Pluie       
                BatteryPercent : 72 (updated 9684499s ago)
                RFStatus : 54 (updated 9684499s ago)       
                Rain : 0.101 (updated 9684499s ago)        
        Module : Living      
                WifiStatus : 37 (updated 278s ago)         
                Temperature : 24 (updated 278s ago)        
                Humidity : 49 (updated 278s ago)           
                CO2 : 733 (updated 278s ago)               
                Noise : 50 (updated 278s ago)              
                Pressure : 1028.1 (updated 278s ago)       
                AbsolutePressure : 1008.4 (updated 278s ago)
```
## Tips
- Only Read() method actually do an API call and refresh all data at once
- Main station is handle as a module, it means that Modules() method returns list of additional modules and station itself.
