# netatmo-api-go
Simple API to access Netatmo weather station data written in Go.

## Quickstart

- [Create a new netatmo app](https://dev.netatmo.com/dev/createapp)
- Generate a new token using the token generator. Scope needed is `read_station`: ![token_generator_netatmo.png](token_generator_netatmo.png)
- Edit ```test/sample.conf```with your credentials 
- run ```go run test/netatmo-api-test.go -f test/sample.conf```
- Output shall look like :
```
Station : Home
        City: Bern
        Country: CH
        Timezone: Europe/Zurich
        Longitude: 7.265078
        Latitude: 46.565312
        Altitude: 540               
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
- Data() returns sensors values (such as temperature) whereas Info() returns module status (such as battery level)
