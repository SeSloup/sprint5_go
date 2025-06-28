package personaldata

import "fmt"

type Personal struct {
	Name   string
	Weight float64
	Height float64
}

func (p Personal) Print() {
	// TODO: реализовать функцию
	fmt.Println(`name: `+p.Name,
		"weight: "+fmt.Sprintf("%f", p.Weight),
		"height: "+fmt.Sprintf("%f", p.Height))
}
