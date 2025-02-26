package main

import "fmt"

const pi = 3.14159

func calcularVolume(R float64) float64 {
	var volume float64
	volume = (4.0/3.0) * pi * (R * R * R)
	return volume
}
func main() {
	var raio float64

	fmt.Scan(&raio)

	fmt.Printf("VOLUME = %.3f\n", calcularVolume(raio))
}
