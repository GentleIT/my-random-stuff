package main

import (
	"fmt"
	"log"
	"math"
)

type Coordinates struct {
	latitude  float32 // Широта
	longitude float32 // Долгота
}

func (s *Coordinates) degreesToRadians() {
	s.latitude = s.latitude * (math.Pi / 180)
	s.longitude = s.longitude * (math.Pi / 180)
}

func main() {
	firstCoordinate := Coordinates{ // Алматы
		latitude:  43.2567,
		longitude: 76.9286,
	}

	secondCoordinate := Coordinates{ // Караганды
		latitude:  49.8019,
		longitude: 73.1021,
	}

	Welcome(&firstCoordinate, &secondCoordinate)

	// Преобразование в радианы
	firstCoordinate.degreesToRadians()
	secondCoordinate.degreesToRadians()

	log.Println(fmt.Sprintf("%.4f%s", MeasureDistanceBetween(firstCoordinate, secondCoordinate), "~")) // Вычисление расстояния: 785.0499 | Geolot.ru показывает: 784.188

	fmt.Scan(&firstCoordinate.latitude) // Для выхода из билда
}

func MeasureDistanceBetween(firstCord Coordinates, secondCord Coordinates) float32 {
	R := 6378.0 // Радиус Земли

	// Haversine formula:
	// a=sin^2(Δϕ/2)+cos(ϕ1)⋅cos(ϕ2)⋅sin^2(Δλ/2)
	sinLatitude := math.Sin((float64(secondCord.latitude) - float64(firstCord.latitude)) / 2) //
	sinLongitude := math.Sin((float64(secondCord.longitude) - float64(firstCord.longitude)) / 2)
	var a float64 = sinLatitude*sinLatitude + math.Cos(float64(firstCord.latitude))*math.Cos(float64(secondCord.latitude))*sinLongitude*sinLongitude

	// c = 2⋅atan2(√a, √1-a)
	var c float64 = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// d = R⋅c
	d := R * c

	return float32(d) // У меня есть мелкие сомнения к преобразованиям в этой функции
}

func Welcome(fc *Coordinates, sc *Coordinates) { // bruuuuuuuh
	log.Println("Enter the first coordinates: (Example: 43.2567 76.9286)")
	fmt.Scan(&fc.latitude, &fc.longitude)
	log.Println("Enter the second coordinates: (Example: 49.8019 73.1021)")
	fmt.Scan(&sc.latitude, &sc.longitude)
}

// 08.04.25 ¯\_(ツ)_/¯
