package space

type Planet string

const earthSeconds = 31557600.0

var orbitalTime = map[Planet]float64{

	"Earth":   earthSeconds,
	"Venus":   earthSeconds * 0.61519726,
	"Mercury": earthSeconds * 0.2408467,
	"Mars":    earthSeconds * 1.8808158,
	"Jupiter": earthSeconds * 11.862615,
	"Saturn":  earthSeconds * 29.447498,
	"Uranus":  earthSeconds * 84.016846,
	"Neptune": earthSeconds * 164.79132}

func Age(seconds float64, planet Planet) float64 {
	return seconds / orbitalTime[planet]
}
