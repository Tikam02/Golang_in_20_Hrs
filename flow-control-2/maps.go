
// Maps

// A map maps keys to values.

// The zero value of a map is nil. A nil map has no keys, nor can keys be added.

// The make function returns a map of the given type, initialized and ready for use.
package main
import "fmt"

type vertex struct {
	Lat, Long float64
}

var m map[string]vertex

func main(){
	m = make(map[string]vertex)
	m["Bell Labs"] = vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}