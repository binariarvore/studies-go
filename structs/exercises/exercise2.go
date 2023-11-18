package exercises

import "fmt"

/*
- Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
- Demonstre os valores do map utilizando range.
- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
*/

func ExerciseTwo() {
	mapPessoa := make(map[string]pessoa)

	mapPessoa["Santos"] = pessoa{"Allison", "Santos", []string{"Uva", "Milho Verde", "Pistache"}}
	mapPessoa["Alves"] = pessoa{"João", "Alves", []string{"Chocolate", "Morango", "Abacaxi"}}

	for _, v := range mapPessoa {
		fmt.Printf("Nome: %v, Sobrenome: %v, Sorvetes favoritos:", v.nome, v.sobrenome)
		for _, s := range v.sorvetesFavoritos {
			fmt.Print(" ", s)
		}
		fmt.Println()
	}
}
