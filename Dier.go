package Animals

import "fmt"

type Dier interface {
	eten() string
	geluidmaken() string
}

func VerzorgDier(dier Dier){
	fmt.Println(dier.eten())
	fmt.Println(dier.geluidmaken())
}