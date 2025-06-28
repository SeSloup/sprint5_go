package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	perdago "github.com/Yandex-Practicum/tracker/internal/personaldata"
	spenergo "github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

type DaySteps struct {
	Personal perdago.Personal
	Steps    int
	Duration time.Duration
}

func (ds *DaySteps) Parse(datastring string) (err error) {

	parts := strings.Split(datastring, ",")
	ds.Steps = 0
	ds.Duration = 0

	if len(parts) != 2 {
		err := fmt.Errorf("incorrect number of parameters.\n expected 3 values.\n actual %d values", len(parts))

		return err
	}
	stepsCount, err := strconv.Atoi(parts[0])

	if err != nil {
		err = fmt.Errorf("error converting steps: %v", err)

		return err
	}

	if stepsCount <= 0 {
		err = fmt.Errorf("error: wrong value for steps. stepsCount = %d", stepsCount)

		return err
	}

	Duration, err := time.ParseDuration(parts[1])
	if err != nil {
		err = fmt.Errorf("error parsing time value: %v.", err)

		return err
	}
	if Duration <= 0 {
		err = fmt.Errorf("error: wrong value for time. Duration = %.2f", Duration)

		return err
	}

	return nil

}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	steps := ds.Steps
	Duration := ds.Duration

	weight := ds.Personal.Weight
	height := ds.Personal.Height

	dist := float64(steps) * stepLength / mInKm

	var ccal float64
	var err error

	ccal, err = spenergo.WalkingSpentCalories(steps, weight, height, Duration)

	return fmt.Sprintf("Количество шагов: %d./n"+
		"Дистанция составила %.2f км./n"+
		"Вы сожгли %.2f ккал./n", steps, dist, ccal), err
}
