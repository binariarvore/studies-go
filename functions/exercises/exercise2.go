package exercises

/*
 Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
- Passe um valor do tipo slice of int como argumento para a função;
*/

func ExerciseTwo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}

	return sum
}
