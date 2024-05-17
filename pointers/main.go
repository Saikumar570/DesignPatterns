// array of pointers

package main

import "fmt"

func main() {
	arr := []int{10, 20, 30}
	var i int
	var ptrArr [3]*int

	for i = 0; i < 3; i++ {
		ptrArr[i] = &arr[i]
	}

	for i = 0; i < 3; i++ {
		fmt.Printf("*ptrArr[%d] = %d\n", i, *ptrArr[i])
	}
}


// double pointer

package main

import "fmt"

func main() {
	var X int = 100
    var ptrA *int = &X

    var ptrB **int = &ptrA
	
	fmt.Printf("Value of variable X: %d\n", X)
    fmt.Printf("Address of variable X: %d\n", &X)
    fmt.Printf("Value of pointer ptrA: %d\n", ptrA)
	fmt.Printf("Address of pointer ptrA: %d\n", &ptrA)
    fmt.Printf("Value of pointer ptrB: %d\n", ptrB)
	
	fmt.Printf("Value of X from ptrA : %d\n", *ptrA)
	fmt.Printf("Value of X from ptrB : %d\n", **ptrB)
}


// pointer to function

package main
import "fmt"

func myfunction(pvar *int){
     *pvar = *pvar + 10
}

func main() {
	var x int = 25
	fmt.Println("Before function call, x =  ", x)

    var intPtr *int = &x 

    myfunction(intPtr)
	fmt.Println("After function call, x =  ", x)

    myfunction(&x)
    fmt.Println("After passing Address, x = ", x)
}

*/



// Pointer to struct

package main

import "fmt"

type Employee struct {
	name string
	empid int
}

func main() {

	emp := Employee{"Raju", 19078}

	pts := &emp

	fmt.Println(pts)

	fmt.Println(pts.name)

	//dereferencing
	fmt.Println((*pts).name)

}




