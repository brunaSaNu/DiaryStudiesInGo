package main

import (
	"fmt"
	"math"
)

func calculaDisctancia(x1 float64, y1 float64, x2 float64, y2 float64) float64{
	var distancia = math.Sqrt(math.Pow((x2 - x1),2) + math.Pow((y2 - y1), 2))
	return distancia
}

func main() {
	var x1, y1, x2, y2 float64
	fmt.Scan(&x1, &y1, &x2, &y2)

	fmt.Printf("%.4f\n", calculaDisctancia(x1,y1,x2,y2))
}