package space

import "math"

type Planet struct {
	Name         string
	OrbialPeriod float64
}

func Age(seconds float64, p Planet) float64 {
	if p.OrbialPeriod < 0 || math.Abs(p.OrbialPeriod) < 0.00001 {
		return math.NaN()
	}
	const secondsPerYearOnEarth float64 = 31557600
	return seconds / secondsPerYearOnEarth / p.OrbialPeriod
}
