# SportCrime [![CircleCI](https://circleci.com/gh/orcaman/spotcrime.svg?style=svg)](https://circleci.com/gh/orcaman/spotcrime)

The following is an unoffical go client for the SpotCrime API. SpotCrime is a website that provides stats about crime in geos.
This unoffical software is not affiliated with spotcrime in any way. The API Spec may be found [here](https://www.programmableweb.com/api/spotcrime).

## Usage Example

Initialize a client, and use the `GetCrimes` function, passing lat/lon coordinates and optionally a seach radius.

```go
package main

import (
	"flag"
	"log"

	"encoding/json"

	"github.com/orcaman/spotcrime"
)

var (
	key    = flag.String("key", "", "spotcrime api key")
	lat    = flag.Float64("lat", 33.39657, "latitude")
	lon    = flag.Float64("lon", -112.03422, "longitude")
	radius = flag.Float64("radius", 0, "search radius")
)

func init() {
	flag.Parse()
}

func main() {
	if len(*key) == 0 {
		log.Fatal("must provide key flag")
	}

	c := spotcrime.New(*key)

	resp, err := c.GetCrimes(&spotcrime.Request{
		Lat:    *lat,
		Lon:    *lon,
		Radius: *radius,
	})

	if err != nil {
		log.Fatalln(err)
	}

	j, err := json.MarshalIndent(resp.Results[0], " ", "\t")

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("response: %s", j)

// this will print for example:    
//     2017/02/01 16:56:16 response: {
//  	"cdid": 88505560,
//  	"type": "Assault",
//  	"date": "01/17/17 11:00 PM",
//  	"link": "http://spotcrime.com/crime/88505560-d0af661686da465e59a16705b03f13e9",
//  	"lat": 33.3969209,
//  	"lon": -112.0344771
//  }
}
```