package main

import (
	"log"
	"math"
	"net/http"
	"strconv"
)

func parseParams(r *http.Request) map[string]float64 {
	params := map[string]float64{
		"width":   600, // canvas size in pixels
		"height":  320,
		"cells":   100,  // number of grid cells
		"xyrange": 30.0, // axis ranges (-xyrange..+xyrange)
	}

	for key, value := range r.URL.Query() {
		// If the key does not already exist in our params map, skip this iteration of the loop and don't process it.
		if _, exists := params[key]; !exists {
			continue
		}

		// Only consider the last value, if multiple are passed in through the URL query parameter.
		fParam, errParse := strconv.ParseFloat(value[len(value)-1], 64)
		if errParse != nil {
			log.Printf("Error parsing float64 from parameter '%s': %v", value, errParse)
			continue // If it doesn't parse correctly, bin it.
		}

		params[key] = fParam
	}

	params["xyscale"] = params["width"] / 2 / params["xyrange"] // pixels per x or y unit
	params["zscale"] = params["height"] * 0.4                   // pixels per z unit
	params["angle"] = math.Pi / 6                               // angle of x, y axes (=30°)

	params["sin30"] = math.Sin(params["angle"]) // sin(30°)
	params["cos30"] = math.Cos(params["angle"]) // cos(30°)

	return params
}
