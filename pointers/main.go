package main

import (
	"fmt"

	"github.com/binariarvore/studies-go/pointers/exercises"
)

func main() {
	y := exercises.ExerciseOne()
	x := &y
	fmt.Println(&x)
	joão := exercises.ExerciseTwo()
	exercises.MudeMe(&joão, "José")
	fmt.Println(joão)
}
