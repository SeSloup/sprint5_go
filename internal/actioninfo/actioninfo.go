package actioninfo

type DataParser interface {
	Parse(dataset string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, tt := range dataset {
		dp.Parse(tt)
		dp.ActionInfo()

	}

}
