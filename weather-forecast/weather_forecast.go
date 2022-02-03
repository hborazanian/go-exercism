// Package weather informs about weather.
package weather

// CurrentCondition informs current condition.
var CurrentCondition string

// CurrentLocation informs current location.
var CurrentLocation string

// Forecast returns sentence with current condition and location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
