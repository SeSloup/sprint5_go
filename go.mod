module github.com/Yandex-Practicum/tracker

go 1.24.3

require github.com/stretchr/testify v1.10.0

require github.com/Yandex-Practicum/tracker/internal/spentenergy v0.0.0-00010101000000-000000000000 // indirect

require (
	github.com/Yandex-Practicum/tracker/internal/daysteps v1.2.3
	github.com/Yandex-Practicum/tracker/internal/personaldata v1.2.3
	github.com/Yandex-Practicum/tracker/internal/trainings v1.2.3
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/Yandex-Practicum/tracker/internal/personaldata => ./internal/personaldata

replace github.com/Yandex-Practicum/tracker/internal/spentenergy => ./internal/spentenergy

replace github.com/Yandex-Practicum/tracker/internal/trainings => ./internal/trainings

replace github.com/Yandex-Practicum/tracker/internal/daysteps => ./internal/daysteps
