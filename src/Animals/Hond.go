package Animals

type Hond struct {
	Name string
}

func (h Hond) eten() string {
	return "Je hebt "+ h.Name +  " eten gegeven"
}

func (h Hond) geluidmaken() string {
	return "Woef woef"
}