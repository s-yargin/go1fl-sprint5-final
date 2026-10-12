package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.ValidateInput
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func validateInputValues(steps int, weight, height float64, duration time.Duration) error {
	switch {
	case steps <= 0:
		return errors.New("steps must be greater than 0")
	case weight <= 0:
		return errors.New("weight must be greater than 0")
	case height <= 0:
		return errors.New("height must be greater than 0")
	case duration <= 0:
		return errors.New("duration must be greater than 0")
	default:
		return nil
	}
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	err := validateInputValues(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return ((weight * meanSpeed * durationInMinutes) / minInH) * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	err := validateInputValues(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return (weight * meanSpeed * durationInMinutes) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	switch {
	case steps <= 0:
		return 0
	case height <= 0:
		return 0
	case duration <= 0:
		return 0
	}
	distance := Distance(steps, height)
	return distance / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	switch {
	case steps <= 0:
		return 0
	case height <= 0:
		return 0
	}
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}
