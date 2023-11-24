package exercises

import (
	"fmt"
	"math"
)

/*
 Crie um tipo "quadrado"
- Crie um tipo "círculo"
- Crie um método "área" para cada tipo que calcule e retorne a área da figura
    - Área do círculo: 2 * π * raio
    - Área do quadrado: lado * lado
- Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
- Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
- Crie um valor de tipo "quadrado"
- Crie um valor de tipo "círculo"
- Use a função "info" para demonstrar a área do "quadrado"
- Use a função "info" para demonstrar a área do "círculo"
*/

type quadrado struct {
	lado float64
}

func (q quadrado) Area() float64 {
	return q.lado * q.lado
}

type circulo struct {
	raio float64
}

func (c circulo) Area() float64 {
	return 2 * math.Pi * c.raio
}

type figura interface {
	Area() float64
}

func info(f figura) string {
	return fmt.Sprintln("A area é: ", f.Area())
}

func ExerciseFive() {
	x := quadrado{15.0}
	y := circulo{5.25}

	fmt.Println(info(x))
	fmt.Println(info(y))
}
