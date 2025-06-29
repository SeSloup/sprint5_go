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
		return fmt.Errorf("incorrect count of parameters.\n expected 2 values.\n actual %d values", len(parts))

	}
	stepsCount, err := strconv.Atoi(parts[0])

	if err != nil {
		return fmt.Errorf("error converting steps: %w", err)

	}

	if stepsCount <= 0 {
		return fmt.Errorf("error: wrong value for steps. stepsCount = %d", stepsCount)

	}

	Duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("error parsing time value: %w.", err)

	}
	if Duration <= 0 {
		return fmt.Errorf("error: wrong value for time. Duration = %.2f", Duration)

	}

	ds.Steps = stepsCount
	ds.Duration = Duration

	return nil

}

func (ds DaySteps) ActionInfo() (string, error) {

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
