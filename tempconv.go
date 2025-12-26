// Package tempconv performs Celsius and Fahrenheit conversions.

package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = 0
	Freezing      Celsius = 0
	BoilingC      Celsius = 0
)

// String returns the Celsius temperature formatted as a string with the °C suffix.
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

// String returns the Fahrenheit temperature formatted as a string with the °F suffix.
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g°K", k)
}
