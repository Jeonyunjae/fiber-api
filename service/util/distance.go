package util

import "math"

// calculate distance by two point
// http://www.movable-type.co.uk/scripts/latlong.html

const (
	// one degree in radians
	radPerDegree = math.Pi / 180.0
	// Earth's radius in kilometres
	earthRadius = 6371
)

// Point represents a single set of coordinates on Earth.
type Point struct {
	lat, lon float64
}

// NewPoint returns a new Point with specified lat/lon coordinates in degrees.
func NewPoint(lat, lon float64) Point {
	return Point{
		lat: lat,
		lon: lon,
	}
}

// RadLat returns point's Latitude in radians.
func (p Point) RadLat() float64 {
	return p.lat * radPerDegree
}

// RadLon returns point's Longtitude in radians.
func (p Point) RadLon() float64 {
	return p.lon * radPerDegree
}

func (p Point) Lat() float64 {
	return p.lat
}

func (p Point) Lon() float64 {
	return p.lon
}

// DistanceHav calculates the distance between two points in kilometres.
// Haversine distance formula is used to calculate the distance.
//func Distance(lon1, lat1, lon2, lat2 float64) float64 {
func distance(lon_P1, lon_P2, lat_P1, lat_P2 float64) float64 {
	// http://www.movable-type.co.uk/scripts/latlong.html

	p1 := Point{lat: lon_P1, lon: lat_P1}
	p2 := Point{lat: lon_P2, lon: lat_P2}
	lat_P1 = p1.RadLat()
	lat_P2 = p2.RadLat()

	deltaLat := lat_P2 - lat_P1
	deltaLon := p2.RadLon() - p1.RadLon()

	sqSinDLat := math.Pow(math.Sin(deltaLat/2), 2)
	sqSinDLon := math.Pow(math.Sin(deltaLon/2), 2)

	// left and right-hand sides of an eq for Haversine
	a := sqSinDLat + sqSinDLon*math.Cos(lat_P1)*math.Cos(lat_P2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}

func GetDistance(lon1, lon2, lat1, lat2 float64) float64 {
	if (lat1 == lat2) && (lon1 == lon2) {
		return 0
	}

	radLat1 := math.Pi * lat1 / 180
	radLat2 := math.Pi * lat2 / 180
	theta := lon1 - lon2
	radTheta := math.Pi * theta / 180
	dist := math.Sin(radLat1)*math.Sin(radLat2) + math.Cos(radLat1)*math.Cos(radLat2)*math.Cos(radTheta)
	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515 * 1.609344 * 1000
	if dist < 100 {
		dist = math.Round(dist/10) * 10
	} else {
		dist = math.Round(dist/100) * 100
	}

	return dist
}
