/*this shows the name of the package the program is in
in this case its in the main package
*/
package main

/*
import statements helps us get external packages to work with in the program
*/
import (
	"fmt"
	/*
		fmt -this package is used for input/output in the program similar to scanf in c but its simpler
		it helps us to print outputs to the terminal
	*/)

/*
main() -this is the part of program that is executed when the program runs.
      -it marks entry point of the program
      -it returns void
	  -
*/
func main() {

	/*
		fmt.Println(...) -this tells the program what to print to the console
	*/
	fmt.Println("this will be printed out on the console")
}
