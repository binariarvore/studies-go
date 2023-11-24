package exercises

/* - Callback: passe uma função como argumento a outra função. */

func Recursive(x int, y []int) int {
	if x == 0 {
		return y[x]
	}

	return y[x] + Recursive(x-1, y)
}

func ExerciseSeven(f func(int, []int) int, x ...int) int {
	slice := []int{}

	for _, v := range x {
		if v%2 == 0 {
			slice = append(slice, v)
		}
	}

	length := len(slice) - 1

	return f(length, slice)
}
