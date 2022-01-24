//Package weather reports the current weather condition at certain location.
package weather

// CurrentCondition is a variable that hosts a string, which describes a certain weather condition.
var CurrentCondition string

// CurrentLocation is a variable that hosts a string, which shows a certain location.
var CurrentLocation string

// Forecast returns a string describing the weather condition at a certain location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
