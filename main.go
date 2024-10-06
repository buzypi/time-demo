package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	_ "time/tzdata" // This embeds timezone data into the binary
)

func main() {
	// Get the TZ environment variable
	tz := os.Getenv("TZ")
	// Handle empty or invalid TZ variable
	if tz == "" {
		fmt.Println("Timezone not set or invalid, falling back to UTC.")
		tz = "UTC"
	}

	// Try to load the location as a named timezone first (e.g., "America/New_York")
	loc, err := time.LoadLocation(tz)
	if err != nil {
		// If it fails, check if it's a fixed UTC offset (e.g., "UTC+5" or "UTC-3")
		if strings.HasPrefix(tz, "UTC") && len(tz) > 3 {
			offsetStr := tz[3:] // Get the "+5" or "-3" part
			offset, err := strconv.Atoi(offsetStr)
			if err == nil {
				// Create a fixed timezone with the specified offset
				loc = time.FixedZone(tz, offset*3600)
			} else {
				// If the offset is invalid, fall back to UTC
				fmt.Println("Invalid UTC offset, falling back to UTC.")
				loc = time.UTC
			}
		} else {
			// If it's not a valid timezone name or offset, fall back to UTC
			fmt.Println("Timezone not set or invalid, falling back to UTC.")
			loc = time.UTC
		}
	}

	// Print the current time in the given timezone
	currentTime := time.Now().In(loc)
	fmt.Printf("Current time in %s: %s\n", tz, currentTime)
}
