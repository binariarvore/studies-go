package exercises

import "fmt"

/*
- Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
    - Nome
    - Sobrenome
    - Sabores favoritos de sorvete
- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
*/

type pessoa struct {
	nome              string
	sobrenome         string
	sorvetesFavoritos []string
}

func ExerciseOne() {
	pessoa1 := pessoa{"Allison", "Santos", []string{"Milho Verde", "Uva"}}

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	for _, v := range pessoa1.sorvetesFavoritos {
		fmt.Println(v)
	}

	pessoa2 := new(pessoa)
	pessoa2.nome = "Mario"
	pessoa2.sobrenome = "Silva"
	pessoa2.sorvetesFavoritos = []string{"Chocolate", "Pistache"}

	fmt.Println(pessoa2.nome, pessoa2.sobrenome)
	for _, v := range pessoa2.sorvetesFavoritos {
		fmt.Println(v)
	}
}
