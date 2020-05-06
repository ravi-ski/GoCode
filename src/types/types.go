package types

import (
	"fmt"
	"os"
)

//Types Demo
func Types() {
	var ( // All variables under var keyword, comprssed declaration
		speed int
		head  float64

		justbool bool
		brand    string
	)

	var spicy, inoneline int

	/* Short Frm declaration */
	var name string = "carl"
	var name1 = "carl2" //Type will inferred from initilization
	name2 := "Carl3"    // Declaration + Initilization
	//when we specify :, no need to mention var word
	// short declaration is not allowed to initilized at package scope
	fmt.Println(name)
	fmt.Println(name1)
	fmt.Println(name2)

	/* Multiple Short Declaration at one shot */

	safe, speed := true, 50
	fmt.Println("Multiple Short Declaration:", safe, speed)

	fmt.Println(speed, head, justbool, brand, spicy, inoneline)

	fmt.Println("==============================================")
	fmt.Println("             RE DECLARATIONS                  ")
	fmt.Println("==============================================")
	fmt.Println("while redeclarations, short declarations one variable should be new variable.")

	var redVar int
	redVar, newVar := 10, 300
	fmt.Println("Redeclarations Types ", redVar, newVar)

	redVar1 := 20 //Declaration + Assignment
	redVar1 = 30  // Assignment
	fmt.Println(redVar1)

	fmt.Println("==============================================")
	fmt.Println("             OS ARGS                  ")
	fmt.Println("==============================================")

	fmt.Println("Total Arguments = ", os.Args)
	fmt.Printf("it's printf O/p = %v\n", os.Args)

	fmt.Println("New Type Declarations are possible using :type gram float64, like this")

}
