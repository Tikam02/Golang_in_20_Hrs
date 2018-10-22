
// Constants

// Constants are declared like variables, but with the const keyword.

// Constants can be character, string, boolean, or numeric values.

// Constants cannot be declared using the := syntax.

package main
import "fmt"

const Pi = 3.1462

func main(){
	const World = "世界"
	fmt.Println("Hello",World)
	fmt.Println("hello",Pi,"Day")

	const Truth = true
	fmt.Println("Go rules?",Truth)
}