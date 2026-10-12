package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	s := strings.Split(datastring, ",")
	if len(s) != 2 {
		err = errors.New("input data must contain 2 fields separated by commas")
		return err
	}

	steps, err := strconv.Atoi(s[0])
	if err != nil {
		return fmt.Errorf("invalid steps format: %w", err)
	}
	if steps <= 0 {
		err = errors.New("steps must be greater than 0")
		return err
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(s[1])
	if err != nil {
		return fmt.Errorf("invalid duration format: %w", err)
	}
	if duration <= 0 {
		err = errors.New("duration must be greater than 0")
		return err
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)
	walkingSpentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, walkingSpentCalories), nil
}
