package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	perdago "github.com/Yandex-Practicum/tracker/internal/personaldata"
	spenergo "github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Personal     perdago.Personal
	Steps        int
	TrainingType string
	Duration     time.Duration
}

func (t *Training) Parse(datastring string) (err error) {

	//"3456,Ходьба,3h00m" - пример корректных входных данных

	parts := strings.Split(datastring, ",")
	t.Steps = 0
	t.TrainingType = ""
	t.Duration = 0

	if len(parts) != 3 {

		return fmt.Errorf("incorrect number of parameters.\n expected 3 values.\n actual %d values", len(parts))
	}

	Steps, err := strconv.Atoi(parts[0])

	if Steps <= 0 {

		return errors.New("error: wrong value for steps. Steps <= 0")
	}

	if err != nil {

		return fmt.Errorf("error converting steps: %w", err)
	}

	Duration, err := time.ParseDuration(parts[2])

	if Duration <= 0 {

		return errors.New("error: wrong value for time. Duration <= 0")
	}
	if err != nil {

		return fmt.Errorf("error parsing time value: %w.", err)
	}

	t.Steps = Steps
	t.TrainingType = parts[1]
	t.Duration = Duration

	return err

}

func (t Training) ActionInfo() (string, error) {

	steps := t.Steps
	activity := t.TrainingType
	Duration := t.Duration

	height := t.Personal.Height
	weight := t.Personal.Weight

	dist := spenergo.Distance(steps, height)
	ms := spenergo.MeanSpeed(steps, height, Duration)

	var ccal float64
	var err error

	switch {
	case activity == "Ходьба":
		ccal, err = spenergo.WalkingSpentCalories(steps, weight, height, Duration)

	case activity == "Бег":
		ccal, err = spenergo.RunningSpentCalories(steps, weight, height, Duration)

	default:
		err := errors.New("неизвестный тип тренировки") //unknown training type
		text := ""

		return text, err
	}

	if err != nil {
		return "", err
	}

	text := fmt.Sprintf("Тип тренировки: %s\n"+
		"Длительность: %.2f ч.\n"+
		"Дистанция: %.2f км.\n"+
		"Скорость: %.2f км/ч\n"+
		"Сожгли калорий: %.2f\n", activity, Duration.Minutes()/60, dist, ms, ccal)

	return text, err

}
