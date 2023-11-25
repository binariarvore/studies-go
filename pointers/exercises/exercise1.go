package exercises

/*
- Crie um valor e atribua-o a uma variável.
- Demonstre o endereço deste valor na memória.
*/

import (
	"math/rand"
)

func ExerciseOne() []int {
	x := make([]int, 0, 10)

	for i := 0; i < cap(x); i++ {
		x = append(x, rand.Int())
	}

	return x
}
