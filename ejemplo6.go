
package main

import "fmt"

func main() {
	fmt.Print("Ingrese una palabra: ")
	var palabra string
	fmt.Scanf("%s", &palabra)
	for i := 0; i < len(palabra); i++{
   	 fmt.Println(palabra[:i+1])
	}
	for i := len(palabra)-1; i > 0; i--{
   	 fmt.Println(palabra[:i])
	}
}


