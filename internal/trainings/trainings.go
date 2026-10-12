package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	s := strings.Split(datastring, ",")
	if len(s) != 3 {
		err = errors.New("input data must contain 3 fields separated by commas")
		return err
	}

	steps, err := strconv.Atoi(s[0])
	switch {
	case err != nil:
		return err
	case steps <= 0:
		err = errors.New("steps must be greater than 0")
		return err
	}
	t.Steps = steps

	t.TrainingType = s[1]

	duration, err := time.ParseDuration(s[2])
	switch {
	case err != nil:
		return err
	case duration <= 0:
		err = errors.New("duration must be greater than 0")
		return err
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)

	switch t.TrainingType {
	case "Бег":
		runningSpentCalories, err := spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			t.TrainingType, t.Duration.Hours(), distance, meanSpeed, runningSpentCalories), nil
	case "Ходьба":
		walkingSpentCalories, err := spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			t.TrainingType, t.Duration.Hours(), distance, meanSpeed, walkingSpentCalories), nil
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
}
