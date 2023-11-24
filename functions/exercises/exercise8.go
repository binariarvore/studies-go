package exercises

/* - Demonstre o funcionamento de um closure.
- Ou seja: crie uma função que retorna outra função, onde esta outra função faz uso de uma variável alem de seu scope.
*/

func ExerciseEight() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
