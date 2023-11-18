package exercises

import "fmt"

/* - Crie e use um struct anônimo.
- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
*/

func ExerciseFour() {
	pessoa := struct {
		nome      string
		idade     int
		cep       map[string]string
		telefones []int
	}{
		nome:      "Allison",
		idade:     24,
		cep:       make(map[string]string),
		telefones: make([]int, 0, 10),
	}
	pessoa.cep["casa"] = "08340530"
	pessoa.telefones = append(pessoa.telefones, 11973905536)

	if _, ok := pessoa.cep["casa"]; ok {
		fmt.Println(pessoa)
	}
}
