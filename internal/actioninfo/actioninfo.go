package actioninfo

import "fmt"

type DataParser interface {
	Parse(dataset string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, tt := range dataset {
		dp.Parse(tt)
		v, err := dp.ActionInfo()

		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(v)

	}

}
