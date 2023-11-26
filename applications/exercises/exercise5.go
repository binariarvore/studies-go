package exercises

import (
	"fmt"
	"sort"
)

/*
- Partindo do código abaixo, ordene os []user por idade e sobrenome.
- https://play.golang.org/p/BVRZTdlUZ_
- Os valores no campo Sayings devem ser ordenados tambem, e demonstrados de maneira esteticamente harmoniosa.
*/

type user5 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ordenarPorIdade []user5

func (o ordenarPorIdade) Len() int           { return len(o) }
func (o ordenarPorIdade) Less(i, j int) bool { return o[i].Age < o[j].Age }
func (o ordenarPorIdade) Swap(i, j int)      { o[i], o[j] = o[j], o[i] }

type ordenarSobrenome []user5

func (o ordenarSobrenome) Len() int           { return len(o) }
func (o ordenarSobrenome) Less(i, j int) bool { return o[i].Last < o[j].Last }
func (o ordenarSobrenome) Swap(i, j int)      { o[i], o[j] = o[j], o[i] }

func ExerciseFive() {
	u1 := user5{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user5{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user5{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user5{u1, u2, u3}

	sort.Sort(ordenarPorIdade(users))
	fmt.Println("Por idade: ")
	formatUser(users)

	sort.Sort(ordenarSobrenome(users))
	fmt.Println("Por sobrenome: ")
	formatUser(users)
}

func formatUser(x []user5) {
	for _, v := range x {
		fmt.Printf("nome: %v, sobrenome: %v, idade: %v \n", v.First, v.Last, v.Age)
		fmt.Println("Citações: ")
		for _, s := range v.Sayings {
			fmt.Println(s)
		}
		fmt.Println()
	}
}
