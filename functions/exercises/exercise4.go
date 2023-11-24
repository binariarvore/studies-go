package exercises

import "fmt"

/* - Crie um tipo struct "pessoa" que contenha os campos:
    - nome
    - sobrenome
    - idade
- Crie um método para "pessoa" que demonstre o nome completo e a idade;
- Crie um valor de tipo "pessoa";
- Utilize o método criado para demonstrar esse valor.
*/

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) Details() string {
	return fmt.Sprintf("%v, %v, %v", p.nome, p.sobrenome, p.idade)
}

var João = pessoa{"João", "Silva", 18}
