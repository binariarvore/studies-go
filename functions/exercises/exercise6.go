package exercises

/*
- Crie uma função que retorna uma função.
- Atribua a função retornada a uma variável.
- Chame a função retornada.
*/

func Calculus(x int) func(int) int {
	return func(y int) int {
		return x * y
	}
}
