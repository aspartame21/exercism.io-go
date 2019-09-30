package space

type Planet string

var orbitalPeriods = map[Planet]float64{
	"Earth":   31557600,
	"Mercury": 7.60054381992e+06,
	"Venus":   1.9414149052176e+07,
	"Mars":    5.935403269008e+07,
	"Jupiter": 3.74355659124e+08,
	"Saturn":  9.292923628848e+08,
	"Uranus":  2.6513700193296e+09,
	"Neptune": 5.200418560032e+09,
}

func Age(seconds float64, planet Planet) float64 {
	return seconds / orbitalPeriods[planet]
}
