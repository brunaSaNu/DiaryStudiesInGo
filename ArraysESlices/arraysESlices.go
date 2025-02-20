package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e slices")

	var array1 [5] string
	array1[0] = "Primeira posição"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	array3 := [...]int{1 , 2, 3, 4, 5}
	fmt.Println(array3)

	slicOne := []int{10, 20, 30, 40, 50}
	
	fmt.Println(slicOne)

	fmt.Println(reflect.TypeOf(slicOne))
	fmt.Println(reflect.TypeOf(array3))

	slicOne = append(slicOne, 20)
	fmt.Println(slicOne)

	sliceTwo := array2[1:4]
	fmt.Println(sliceTwo)

	array2[1] = "New position"
	fmt.Println(sliceTwo)
}