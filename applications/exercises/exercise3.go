package exercises

/*- Partindo do código abaixo, utilize NewEncoder() e Encode() para enviar as informações como JSON para Stdout.*/

import (
	"encoding/json"
	"os"
)

type user3 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func ExerciseThree() {
	u1 := user3{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user3{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user3{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user3{u1, u2, u3}

	json.NewEncoder(os.Stdout).Encode(users)
}
