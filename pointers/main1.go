// pointers usage in the interfaces

package main

import (
	"fmt"
)

type age interface {
	show(int)
}

func show(a age, newage int) {
	a.show(newage)
}

type ageofone struct {
	oldage int
}

func (aoo ageofone) show(newage int) {
	aoo.oldage = newage
}

type ageofsecond struct {
	oldage int
}

func (aos *ageofsecond) show(newage int) {
	aos.oldage = newage
}
func main() {
	first := ageofone{oldage: 21}
	second := ageofsecond{oldage: 30}
	//after four years updating the ages of two persons
	//using value
	show(first, 25)
	//using reference
	show(&second, 34)
	fmt.Println("The age of first one is ", first.oldage)//The age wont be updated as we are passing by the value
	fmt.Println("The age of second one is ", second.oldage)//The age will be updated to 34 as we are using reference
}
