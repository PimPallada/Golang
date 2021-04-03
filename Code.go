package main
import (
    "Animals"
	"Test"
	"fmt"
)
func main(){
	kat := Animals.Kat{"Klaas"}
	hond := Animals.Hond{"Kees"}
	Animals.VerzorgDier(kat)
	Animals.VerzorgDier(hond)
    fmt.Println(Test.Keer(5,6))
}