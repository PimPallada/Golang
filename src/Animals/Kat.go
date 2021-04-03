package Animals
import ( 
	"Test"
	"fmt"
)

type Kat struct {
	Name string
}


func (k Kat) eten() string {
	fmt.Println(Test.Keer(12,3))
	return "Je hebt " +  k.Name +" eten gegeven"
}

func (k  Kat) geluidmaken() string{
	return "Miauw"
}