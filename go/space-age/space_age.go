// Package space calculates age for given planet
package space

// Planet represents planet in string
type Planet string

// Ages mapping between planet name and time
var Ages = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age calculates age for given time and planet name
func Age(seconds float64, planet Planet) float64 {
	return seconds / 31557600 / Ages[planet]
}
