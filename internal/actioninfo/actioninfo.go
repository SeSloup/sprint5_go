package actioninfo

import "fmt"

type DataParser interface {
	Parse(dataset string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {

	for _, tt := range dataset {
		err := dp.Parse(tt)

		if err != nil {
			fmt.Println("error parse:", err)
			continue
		}

		v, err := dp.ActionInfo()

		if err != nil {
			fmt.Println("error info:", err)
			continue
		}
		fmt.Println(v)

	}

}
