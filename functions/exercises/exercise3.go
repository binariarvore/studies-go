package exercises

import "fmt"

/* - Utilize a declaração defer de maneira que demonstre que sua execução só ocorre ao final do contexto ao qual ela pertence. */

func ExerciseThree() {
	fmt.Println("Primeira mensagem")
	fmt.Println("Segunda mensagem")
	defer fmt.Println("Quinta mensagem")
	defer fmt.Println("Quarta mensagem")
	fmt.Println("Terceira mensagem")
}
